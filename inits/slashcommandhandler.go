package inits

import (
	"archroid/ElProfessorBot/static"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekrotja/ken"
)

func InitCommandHandler(container di.Container) (k *ken.Ken, err error) {
	session := container.Get(static.DiDiscordSession).(*discordgo.Session)

	k, err = ken.New(session, ken.Options{
		DependencyProvider: container,
		OnSystemError:      systemErrorHandler,
		OnCommandError:     commandErrorHandler,
	})

	if err != nil {
		return
	}

	err = k.RegisterCommands()
	if err != nil {
		return
	}

	return
}

func systemErrorHandler(context string, err error, args ...interface{}) {
	logrus.WithField("ctx", context).WithError(err).Error("ken error")
}

func commandErrorHandler(err error, ctx *ken.Ctx) {
	// Is ignored if interaction has already been responded
	ctx.Defer()

	if err == ken.ErrNotDMCapable {
		ctx.FollowUpError("This command can not be used in DMs.", "")
		return
	}

	ctx.FollowUpError(
		fmt.Sprintf("The command execution failed unexpectedly:\n```\n%s\n```", err.Error()),
		"Command execution failed")
}
