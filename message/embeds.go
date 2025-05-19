package message

import "github.com/bwmarrin/discordgo"

// SimpleEmbed creates a basic embed with title, description, and color.
func SimpleEmbed(title, description string, color int) *discordgo.MessageEmbed {
    return &discordgo.MessageEmbed{
        Title:       title,
        Description: description,
        Color:       color,
    }
}
