package aternos_discord_bot

import (
    "context"
    "fmt"
    "strings"

    "github.com/bwmarrin/discordgo"
    aternos "github.com/sleeyax/aternos-api"
    "github.com/sleeyax/aternos-discord-bot/database"
    "github.com/sleeyax/aternos-discord-bot/database/models"
    "github.com/sleeyax/aternos-discord-bot/message"
)

// handleCommands responds to incoming interactive commands on Discord. func (ab *Bot) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) { command := i.ApplicationCommandData()

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
		sendErrorText("Failed to save configuration.", err)
		break
	}
	sendText(message.FormatSuccess("Configuration changed successfully."))

case StatusCommand, InfoCommand, PlayersCommand, StartCommand:
	w, err := ab.getWorker(i.GuildID)
	if err != nil {
		if err == database.ErrDataNotFound {
			sendText(message.FormatWarning("Bot setup incomplete. Use `/configure` to configure the bot."))
		} else {
			sendErrorText("Failed to get worker", err)
		}
		break
	}

	serverInfo, err := w.GetServerInfo()
	if err != nil {
		if err == aternos.UnauthenticatedError {
			sendText(message.FormatError("Invalid credentials. Use `/configure` to reconfigure the bot."))
		} else if err == aternos.ForbiddenError {
			sendText(message.FormatError("Access forbidden. Please try again later."))
		} else {
			sendErrorText("Failed to get server info", err)
		}
		break
	}

	switch command.Name {
	case InfoCommand:
		respondWithEmbeds(s, i, []*discordgo.MessageEmbed{
			message.CreateServerInfoEmbed(serverInfo),
		})

	case StatusCommand:
		embed := message.SimpleEmbed("Server Status", fmt.Sprintf("Server **%s** is currently **%s**.", serverInfo.Name, serverInfo.StatusLabel), message.Blue)
		respondWithEmbeds(s, i, []*discordgo.MessageEmbed{embed})

	case PlayersCommand:
		var embed *discordgo.MessageEmbed
		if len(serverInfo.PlayerList) == 0 {
			embed = message.SimpleEmbed("Online Players", "No players online right now.", message.Yellow)
		} else {
			embed = message.SimpleEmbed("Online Players", fmt.Sprintf("Active players (%d):\n`%s`", len(serverInfo.PlayerList), strings.Join(serverInfo.PlayerList, "`, `")), message.Green)
		}
		respondWithEmbeds(s, i, []*discordgo.MessageEmbed{embed})

	case StartCommand:
		if err = w.Init(); err != nil {
			sendErrorText("Failed to initialize worker! See `/help` or try again later.", err)
			break
		}

		if serverInfo.Status != aternos.Offline && serverInfo.Status != aternos.Stopping {
			embed := message.SimpleEmbed("Server Already Running", "Server is already started! Use `/status` or `/info` to check it.", message.Orange)
			respondWithEmbeds(s, i, []*discordgo.MessageEmbed{embed})
			break
		}

		startingEmbed := message.SimpleEmbed("Starting Server", "The server is booting up. Sit tight!", message.Aqua)
		respondWithEmbeds(s, i, []*discordgo.MessageEmbed{startingEmbed})

		ctx, cancel := context.WithCancel(context.Background())
		go w.On(ctx, func(messageType string, info *aternos.ServerInfo) {
			switch messageType {
			case "ready":
				if command.Name == StartCommand {
					if err = w.Start(); err != nil {
						embed := message.SimpleEmbed("Start Failed", "Failed to start the server. Reconfigure the bot with `/configure` and try again.", message.Red)
						s.ChannelMessageSendEmbed(i.ChannelID, embed)
						w.Log(err.Error())
						cancel()
						return
					}
				}
			case "status":
				if info.Status == aternos.Offline || info.Status == aternos.Online {
					notification, _ := message.CreateServerStatusNotificationEmbed(info)
					s.ChannelMessageSendEmbed(i.ChannelID, notification)
				} else if info.Status == aternos.Preparing {
					embed := message.SimpleEmbed("Queueing...", fmt.Sprintf("Waiting in queue: **%d/%d**, ETA: **%s**", info.Queue.Position, info.Queue.Count, info.Queue.Time), message.Yellow)
					s.ChannelMessageSendEmbed(i.ChannelID, embed)
					w.Log("Waiting in queue...")
				}
			case "connection_error":
				embed := message.SimpleEmbed("Connection Error", "Failed to initialize worker (websocket timeout). Try again later.", message.Red)
				s.ChannelMessageSendEmbed(i.ChannelID, embed)
			}
		})
	}
default:
	sendText(message.FormatWarning("Command unavailable. Please try again later or refresh your Discord client with `CTRL + R`."))
}

}

