package message

import "github.com/bwmarrin/discordgo"

// CreateHelpEmbed returns a Discord embed for the /help command.
func CreateHelpEmbed() *discordgo.MessageEmbed {
        return &discordgo.MessageEmbed{
                Title:       "Need Help?",
                Description: "Thanks for using the Minecraft Server Bot! Here's some useful info:",
                Color:       0x5865F2, // Discord blurple
                Fields: []*discordgo.MessageEmbedField{
                        {
                                Name: "Why can't I start my new server?",
                                Value: "Go to [aternos.org](https://aternos.org), start your server manually **once**, and accept the EULA. After that, you can use this bot to start it.",
                        },
                        {
                                Name: "The bot suddenly stopped working?",
                                Value: "You were probably logged out of Aternos. Use `/configure` again with a new session and server cookie.",
                        },
                        {
                                Name: "Need more help?",
                                Value: "[GitHub repo](https://github.com/epicvasia/aternos-discord-bot) – Open an issue or discussion if you're stuck.",
                        },
                },
                Footer: &discordgo.MessageEmbedFooter{
                        Text: "Happy hosting!",
                },
        }
}
