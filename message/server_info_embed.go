package message

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)


func CreateServerInfoEmbed(info *aternos.ServerInfo) *discordgo.MessageEmbed {
	if info.DynIP == "" {
		info.DynIP = "Unavailable"
	}

	return &discordgo.MessageEmbed{
		Title:       "Minecraft Server Info",
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
			Text: "Minecraft Server Bot",
		},
	}
}