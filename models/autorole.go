package models

type AutoRole struct {
	RoleID  string
	GuildID string
}

type WelcomeMessage struct {
	WelcomeChannelId string
	WelcomeMessage   string
	GuildId          string
}
