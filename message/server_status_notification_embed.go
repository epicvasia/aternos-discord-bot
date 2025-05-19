package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)

func CreateServerStatusNotificationEmbed(info *aternos.ServerInfo) (*discordgo.MessageEmbed, error) {
	switch info.Status {
	case aternos.Online:
		return &discordgo.MessageEmbed{
			Title:       "<:online:1373961339432730694> Server is online",
			Description: fmt.Sprintf("Join now! Only %d seconds left.", info.Countdown),
			Color:       colorMap[aternos.Online],
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Server address",
					Value: fmt.Sprintf("`%s`", info.Address),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Dyn IP",
					Value:  fmt.Sprintf("`%s`", info.DynIP),
					Inline: true,
				},
			},
		}, nil
	case aternos.Offline:
		return &discordgo.MessageEmbed{
			Title:       "<:offline:1373961374148988969> Server is offline",
			Description: "The server is currently offline.",
			Color:       colorMap[aternos.Offline],
		}, nil
	default:
		return nil, fmt.Errorf("unknown server status code '%d'", info.Status)
	}
}
