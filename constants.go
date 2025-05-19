package aternos_discord_bot

import (
	"github.com/bwmarrin/discordgo"
)

func CreateHelpEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Minecraft Server Bot Help",
		Description: "Thanks for using the Minecraft Server Bot! Here's a quick FAQ to help you out:",
		Color:       0x5865F2, // Discord blurple
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Why can't I start my new server?",
				Value: "Go to [aternos.org](https://aternos.org), start your server manually once, and accept the EULA. After that, you can start it using the bot!",
			},
			{
				Name: "Why did the bot stop working after I logged out?",
				Value: "This is expected. Run `/configure` again to re-authenticate. The bot doesn't use username:password login (and likely never will).",
			},
			{
				Name: "Need help or found a bug?",
				Value: "Check out the GitHub repo: [github.com/epicvasia/aternos-discord-bot](https://github.com/epicvasia/aternos-discord-bot). You can open an issue or discussion there!",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Happy hosting!",
		},
	}
}