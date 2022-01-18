package inits

import (
	"archroid/ElProfessorBot/listeners"
	"archroid/ElProfessorBot/static"
	"os"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/DisgoOrg/disgolink/lavalink"
	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	Link *dgolink.Link
}

func InitDiscordBotSession(container di.Container) *dgolink.Link {

	session := container.Get(static.DiDiscordSession).(*discordgo.Session)

	session.Token = "Bot " + os.Getenv("DISCORD_BOT_TOKEN")

	session.Identify.Intents = discordgo.MakeIntent(static.Intents)

	bot := &Bot{
		Link: dgolink.New(session),
	}

	session.AddHandler(listeners.NewListenerReady(container).Handler)
	session.AddHandler(listeners.NewListenerMemberAdd(container).Handler)

	err := session.Open()
	if err != nil {
		logrus.WithError(err).Fatal("Failed connecting Discord bot session")
	}

	bot.Link.AddNode(lavalink.NodeConfig{
		Name:     "test",
		Host:     "localhost",
		Port:     "2333",
		Password: "1274",
	})
	return bot.Link
}
