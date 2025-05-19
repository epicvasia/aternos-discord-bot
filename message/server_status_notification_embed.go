package message

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)

func CreateServerStatusNotificationEmbed(info *aternos.ServerInfo) (*discordgo.MessageEmbed, error) {
	switch info.Status {
	case aternos.Online:
		dynIP := info.DynIP
		if dynIP == "" {
			dynIP = "Unavailable"
		}

		return &discordgo.MessageEmbed{
			Title:       "<:online:1373961339432730694> Server Online",
			Description: fmt.Sprintf("**%s** is now online! Join before it shuts down in **%d seconds**.", info.Name, info.Countdown),
			Color:       colorMap[aternos.Online],
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Server Address",
					Value:  fmt.Sprintf("`%s`", info.Address),
					Inline: true,
				},
				{
					Name:   "Dyn IP",
					Value:  fmt.Sprintf("`%s`", dynIP),
					Inline: true,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Minecraft Server",
			},
		}, nil

	case aternos.Offline:
		return &discordgo.MessageEmbed{
			Title:       "<:offline:1373961374148988969> Server Offline",
			Description: fmt.Sprintf("**%s** is currently shut down.", info.Name),
			Color:       colorMap[aternos.Offline],
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Minecraft Server",
			},
		}, nil

	default:
		return nil, fmt.Errorf("unknown server status code '%d'", info.Status)
	}
}