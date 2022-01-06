package inits

import (
	"archroid/ElProfessorBot/listeners"
	"archroid/ElProfessorBot/static"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
)

func InitDiscordBotSession(container di.Container) {

	session := container.Get(static.DiDiscordSession).(*discordgo.Session)

	session.Token = "Bot " + os.Getenv("DISCORD_BOT_TOKEN")

	session.StateEnabled = true
	session.Identify.Intents = discordgo.MakeIntent(static.Intents)
	session.StateEnabled = false

	session.AddHandler(listeners.NewListenerReady(container).Handler)
	session.AddHandler(listeners.NewListenerMemberAdd(container).Handler)

	err := session.Open()
	if err != nil {
		logrus.WithError(err).Fatal("Failed connecting Discord bot session")
	}

}
