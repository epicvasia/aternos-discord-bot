package aternos_discord_bot

import "github.com/bwmarrin/discordgo"

const (
	HelpCommand      = "help"
	PingCommand      = "ping"
	ConfigureCommand = "configure"
	StartCommand     = "start"
	StatusCommand    = "status"
	InfoCommand      = "info"
	PlayersCommand   = "players"
	SessionOption    = "session"
	ServerOption     = "server"
)

var (
	adminPermissions int64 = discordgo.PermissionManageServer
	userPermissions  int64 = discordgo.PermissionUseSlashCommands
	dmPermission           = false
)

// List of available discord commands.
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        ConfigureCommand,
		Description: "Save configuration settings",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         SessionOption,
				Description:  "Set the ATERNOS_SESSION cookie value",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildText},
			},
			{
				Name:         ServerOption,
				Description:  "Set the ATERNOS_SERVER cookie value",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildText},
			},
		},
		DefaultMemberPermissions: &adminPermissions, // keep admin only
		DMPermission:             &dmPermission,
	},
	{
		Name:        StartCommand,
		Description: "Start the minecraft server",
		// No admin restriction here
		DMPermission: &dmPermission,
	},
	{
		Name:        PingCommand,
		Description: "Check if the discord bot is still alive",
		DMPermission: &dmPermission,
	},
	{
		Name:        StatusCommand,
		Description: "Get the minecraft server status",
		DMPermission: &dmPermission,
	},
	{
		Name:        InfoCommand,
		Description: "Get detailed information about the minecraft server status",
		DMPermission: &dmPermission,
	},
	{
		Name:        PlayersCommand,
		Description: "List active players",
		DMPermission: &dmPermission,
	},
	{
		Name:        HelpCommand,
		Description: "Get help",
		DMPermission: &dmPermission,
	},
}
