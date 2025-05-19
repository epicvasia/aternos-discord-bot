package message

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)

var colorMap = map[aternos.ServerStatus]int{
	aternos.Online:    0x57F287, // green
	aternos.Offline:   0xED4245, // red
	aternos.Starting:  0xFEE75C, // yellow
	aternos.Stopping:  0xFEE75C, // yellow
	aternos.Loading:   0x5865F2, // blurple
	aternos.Preparing: 0x5865F2, // blurple
	aternos.Saving:    0x979C9F, // gray
}

func CreateServerInfoEmbed(info *aternos.ServerInfo) *discordgo.MessageEmbed {
	if info.DynIP == "" {
		info.DynIP = "Unavailable"
	}

	return &discordgo.MessageEmbed{
		Title:       "Aternos Server Info",
		Description: fmt.Sprintf("**%s** is currently **%s**.", info.Name, info.StatusLabel),
		Color:       colorMap[info.Status],
		URL:         "https://aternos.org/server/",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Players Online",
				Value:  fmt.Sprintf("%d / %d", info.Players, info.MaxPlayers),
				Inline: true,
			},
			{
				Name:   "Problems",
				Value:  fmt.Sprintf("%d", info.Problems),
				Inline: true,
			},
			{
				Name:   "Software",
				Value:  fmt.Sprintf("%s v%s", info.Software, info.Version),
				Inline: true,
			},
			{
				Name:   "Server Address",
				Value:  fmt.Sprintf("`%s`", info.Address),
				Inline: true,
			},
			{
				Name:   "Dyn IP",
				Value:  fmt.Sprintf("`%s`", info.DynIP),
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Aternos Discord Bot",
		},
	}
}