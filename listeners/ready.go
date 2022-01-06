package listeners

import (
	"archroid/ElProfessorBot/static"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/zekrotja/dgrs"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListenerReady struct {
	db *mongo.Database
	st *dgrs.State
}

func NewListenerReady(container di.Container) *ListenerReady {
	return &ListenerReady{
		db: container.Get(static.DiDatabase).(*mongo.Database),
		st: container.Get(static.DiState).(*dgrs.State),
	}
}

func (l *ListenerReady) Handler(s *discordgo.Session, e *discordgo.Ready) {

}
