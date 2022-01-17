package static

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

const (
	DiContainer            = "container"
	DiDatabase             = "database"
	DiDiscordSession       = "discordgosession"
	DiCommandHandler       = "kencommandhandler"
	DiYoutubeSearch        = "youtubeseach"
	DiLegacyCommandHandler = "skcommandhandler"
	DiDgoLink              = "dgolink"

	DiscordInviteLink = "https://discord.com/oauth2/authorize?client_id=901356147720749096&permissions=2080374975&scope=bot+applications.commands+identify+guilds"

	// DiCommandPrefix = "commandprefix"

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
	ColorEmbedDefault = 0x372168
	ColorEmbedUpdated = 0x8bc34a
	ColorEmbedGray    = 0xb0bec5
	ColorEmbedOrange  = 0xfb8c00
	ColorEmbedGreen   = 0x8Bff4A
	ColorEmbedCyan    = 0x00BCD4
	ColorEmbedYellow  = 0xFFC107
	ColorEmbedViolett = 0x6A1B9A
)

var (
	UrlPattern = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
)
