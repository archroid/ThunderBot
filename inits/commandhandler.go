package inits

import (
	"archroid/ElProfessorBot/commands"
	"archroid/ElProfessorBot/static"
	"archroid/ElProfessorBot/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekroTJA/shireikan"
)

func InitLegacyCommandHandler(container di.Container) shireikan.Handler {
	session := container.Get(static.DiDiscordSession).(*discordgo.Session)

	cmdHandler := shireikan.New(&shireikan.Config{
		GeneralPrefix:         "-",
		AllowBots:             false,
		AllowDM:               false,
		DeleteMessageAfter:    true,
		ExecuteOnEdit:         true,
		InvokeToLower:         true,
		UseDefaultHelpCommand: true,
		
		ObjectContainer: container,
		OnError:         legacyErrorHandler,
	})
	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterCommand(&commands.CmdInfo{})
	cmdHandler.RegisterCommand(&commands.CmdClear{})
	cmdHandler.RegisterCommand(&commands.CmdPlay{})

	logrus.WithField("n", len(cmdHandler.GetCommandInstances())).Info("Commands registered")

	cmdHandler.Setup(session)

	return cmdHandler
}

func legacyErrorHandler(ctx shireikan.Context, errTyp shireikan.ErrorType, err error) {
	switch errTyp {

	// Command execution failed
	case shireikan.ErrTypCommandExec:
		msg, _ := ctx.ReplyEmbedError(
			fmt.Sprintf("Command execution failed unexpectedly: ```\n%s\n```", err.Error()),
			"Command Execution Failed")
		utils.DeleteMessageLater(ctx.GetSession(), msg, 60*time.Second)

	// Failed getting channel
	case shireikan.ErrTypGetChannel:
		msg, _ := ctx.ReplyEmbedError(
			fmt.Sprintf("Failed getting channel: ```\n%s\n```", err.Error()),
			"Unexpected Error")
		utils.DeleteMessageLater(ctx.GetSession(), msg, 60*time.Second)

	// Failed getting channel
	case shireikan.ErrTypGetGuild:
		msg, _ := ctx.ReplyEmbedError(
			fmt.Sprintf("Failed getting guild: ```\n%s\n```", err.Error()),
			"Unexpected Error")
		utils.DeleteMessageLater(ctx.GetSession(), msg, 60*time.Second)

	// Middleware failed
	case shireikan.ErrTypMiddleware:
		msg, _ := ctx.ReplyEmbedError(
			fmt.Sprintf("Command Handler Middleware failed: ```\n%s\n```", err.Error()),
			"Unexpected Error")
		utils.DeleteMessageLater(ctx.GetSession(), msg, 60*time.Second)

	// Middleware failed
	case shireikan.ErrTypNotExecutableInDM:
		msg, _ := ctx.ReplyEmbedError(
			"This command is not executable in DM channels.", "")
		utils.DeleteMessageLater(ctx.GetSession(), msg, 8*time.Second)

	// Ignored Errors
	case shireikan.ErrTypCommandNotFound, shireikan.ErrTypDeleteCommandMessage:
		return
	}
}
