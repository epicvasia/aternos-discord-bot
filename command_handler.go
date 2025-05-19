package aternos_discord_bot

import (
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
	"github.com/sleeyax/aternos-discord-bot/database"
	"github.com/sleeyax/aternos-discord-bot/database/models"
	"github.com/sleeyax/aternos-discord-bot/message"
)

// handleCommands responds to incoming interactive commands on Discord.
func (ab *Bot) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()

	// wrap functions around our utilities to make life easier
	sendText := func(content string) {
		respondWithText(s, i, content)
	}
	sendHiddenText := func(content string) {
		respondWithHiddenText(s, i, content)
	}
	sendErrorText := func(content string, err error) {
		respondWithError(s, i, content, err)
	}

	switch command.Name {
	case HelpCommand:
		respondWithEmbeds(s, i, []*discordgo.MessageEmbed{
			message.CreateHelpEmbed(),
		})

	case PingCommand:
		sendHiddenText(message.FormatDefault("Pong!"))

	case ConfigureCommand:
		options := optionsToMap(command.Options)

		err := ab.Database.UpdateServerSettings(&models.ServerSettings{
			GuildID:       i.GuildID,
			SessionCookie: options[SessionOption].StringValue(),
			ServerCookie:  options[ServerOption].StringValue(),
		})
		if err != nil {
			sendErrorText("❌ Failed to save configuration.", err)
			break
		}
		sendText(message.FormatSuccess("✅ Configuration changed successfully."))

	case StatusCommand, InfoCommand, PlayersCommand, StartCommand:
		w, err := ab.getWorker(i.GuildID)
		if err != nil {
			if err == database.ErrDataNotFound {
				sendText(message.FormatWarning("⚠️ Bot setup incomplete. Use `/configure` to configure the bot."))
			} else {
				sendErrorText("❌ Failed to get worker", err)
			}
			break
		}

		serverInfo, err := w.GetServerInfo()
		if err != nil {
			if err == aternos.UnauthenticatedError {
				sendText(message.FormatError("❌ Invalid credentials. Use `/configure` to reconfigure the bot."))
			} else if err == aternos.ForbiddenError {
				sendText(message.FormatError("❌ Access forbidden. Please try again later."))
			} else {
				sendErrorText("⚠️ Failed to get server info", err)
			}
			break
		}

		switch command.Name {
		case InfoCommand:
			respondWithEmbeds(s, i, []*discordgo.MessageEmbed{
				message.CreateServerInfoEmbed(serverInfo),
			})

		case StatusCommand:
			sendText(message.FormatInfo("Server '%s' is currently **%s**.", serverInfo.Name, serverInfo.StatusLabel))

		case PlayersCommand:
			if len(serverInfo.PlayerList) == 0 {
				sendText(message.FormatInfo("No players online right now."))
				break
			}
			sendText(message.FormatInfo("Active players: %s.", strings.Join(serverInfo.PlayerList, ", ")))

		case StartCommand:
			if err = w.Init(); err != nil {
				sendErrorText("Failed to initialize worker! See `/help` or try again later.", err)
				break
			}

			if serverInfo.Status != aternos.Offline && serverInfo.Status != aternos.Stopping {
				sendText(message.FormatInfo("Server already started! Type `/status` or `/info` to view the status."))
				break
			}

			sendText(message.FormatInfo("Starting the server, please wait..."))

			ctx, cancel := context.WithCancel(context.Background())

			go w.On(ctx, func(messageType string, info *aternos.ServerInfo) {
				switch messageType {
				case "ready":
					if command.Name == StartCommand {
						if err = w.Start(); err != nil {
							s.ChannelMessageSend(i.ChannelID, message.FormatError("Failed to start! Reconfigure the bot with `/configure` and try again. See `/help` if the problem persists."))
							w.Log(err.Error())
							cancel()
							break
						}
					}
				case "status":
					if info.Status == aternos.Offline || info.Status == aternos.Online {
						notification, _ := message.CreateServerStatusNotificationEmbed(info)
						s.ChannelMessageSendEmbed(i.ChannelID, notification)
					} else if info.Status == aternos.Preparing {
						s.ChannelMessageSend(i.ChannelID, message.FormatInfo("Waiting in queue (%d/%d, %s)...", info.Queue.Position, info.Queue.Count, info.Queue.Time))
						w.Log("Waiting in queue...")
					}
				case "connection_error":
					s.ChannelMessageSend(i.ChannelID, message.FormatError("Failed to initialize worker (websocket connection timeout)! See `/help` or try again later."))
				}
			})
		}
	default:
		sendText(message.FormatWarning("Command unavailable. Please try again later or refresh your Discord client with `CTRL + R`."))
	}
}