package inits

import (
	"archroid/ElProfessorBot/static"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/zekrotja/dgrs"
)

func InitState(container di.Container) (s *dgrs.State, err error) {
	session := container.Get(static.DiDiscordSession).(*discordgo.Session)

	return dgrs.New(dgrs.Options{
		DiscordSession: session,
		FetchAndStore:  true,
		Lifetimes: dgrs.Lifetimes{
			Message: 14 * 24 * time.Hour, // 14 Days
		},
	})
}
