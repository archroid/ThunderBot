package listeners

import (
	"archroid/ElProfessorBot/static"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListenerReady struct {
	db *mongo.Database
}

func NewListenerReady(container di.Container) *ListenerReady {
	return &ListenerReady{
		db: container.Get(static.DiDatabase).(*mongo.Database),
	}
}

func (l *ListenerReady) Handler(s *discordgo.Session, e *discordgo.Ready) {

	usd := discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "-help",
				Type: 3,
			},
		},
	}

	s.UpdateStatusComplex(usd)
}
