package listeners

import (
	"archroid/ElProfessorBot/models"
	"archroid/ElProfessorBot/pkg/embedbuilder"
	"archroid/ElProfessorBot/static"
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListenerMemberAdd struct {
	db *mongo.Database
}

func NewListenerMemberAdd(container di.Container) *ListenerMemberAdd {
	return &ListenerMemberAdd{
		db: container.Get(static.DiDatabase).(*mongo.Database),
	}
}

func (l *ListenerMemberAdd) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	filter := bson.M{"guildid": e.GuildID}

	var welcomeMessage models.WelcomeMessage
	err := l.db.Collection("welcome").FindOne(context.TODO(), filter).Decode(&welcomeMessage)
	if err != nil {
		logrus.WithError(err).WithField("gid", e.GuildID).Error("Failed updating welcome message settings")
	}
	txt := ""
	if strings.Contains(welcomeMessage.WelcomeMessage, "[ment]") {
		txt = e.User.Mention()
	}
	msg := strings.Replace(welcomeMessage.WelcomeMessage, "[user]", e.User.Username, -1)
	msg = strings.Replace(msg, "[ment]", e.User.Mention(), -1)

	s.ChannelMessageSendComplex(welcomeMessage.WelcomeChannelId, &discordgo.MessageSend{
		Content: txt,
		Embed: embedbuilder.New().
			WithColor(static.ColorEmbedDefault).
			WithDescription(msg).
			Build(),
	})

	var autoRole models.AutoRole

	err = l.db.Collection("auto-role").FindOne(context.TODO(), filter).Decode(&autoRole)
	if err != nil {
		logrus.WithError(err).WithField("gid", e.GuildID).Error("Failed getting guild autorole from database")
	}
	err = s.GuildMemberRoleAdd(e.GuildID, e.User.ID, autoRole.RoleID)
	if err != nil {
		logrus.WithError(err).WithField("gid", e.GuildID).WithField("uid", e.User.ID).Error("Failed setting autorole for member")
	}

}
