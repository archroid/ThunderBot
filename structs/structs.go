package structs

type WelcomeMessage struct {
	WelcomeChannelId string
	WelcomeMessage   string
	GuildId          string
}

type Role struct {
	RoleID  string
	GuildID string
}

type Rules struct {
	Rules   string
	GuildID string
}
