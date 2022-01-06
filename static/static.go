package static

import "github.com/bwmarrin/discordgo"

const (
	DiDatabase       = "database"
	DiDiscordSession = "discordgosession"
	// DiConfig         = "config"
	DiState          = "dgstate"
	DiCommandHandler = "kencommandhandler"

	Intents = discordgo.IntentsDirectMessages |
		discordgo.IntentsGuildBans |
		discordgo.IntentsGuildEmojis |
		discordgo.IntentsGuildIntegrations |
		discordgo.IntentsGuildInvites |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildVoiceStates |
		discordgo.IntentsGuilds

	ColorEmbedError   = 0xd32f2f
	ColorEmbedDefault = 0xffc107
	ColorEmbedUpdated = 0x8bc34a
	ColorEmbedGray    = 0xb0bec5
	ColorEmbedOrange  = 0xfb8c00
	ColorEmbedGreen   = 0x8BC34A
	ColorEmbedCyan    = 0x00BCD4
	ColorEmbedYellow  = 0xFFC107
	ColorEmbedViolett = 0x6A1B9A
)
