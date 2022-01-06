package main

import (
	"archroid/ElProfessorBot/inits"
	"archroid/ElProfessorBot/static"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekrotja/ken"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	// Initialize dependency injection builder
	diBuilder, _ := di.NewBuilder()

	// Initialize discord bot session and shutdown routine
	diBuilder.Add(di.Def{
		Name: static.DiDiscordSession,
		Build: func(ctn di.Container) (interface{}, error) {
			return discordgo.New()
		},
		Close: func(obj interface{}) error {
			session := obj.(*discordgo.Session)
			logrus.Info("Shutting down bot session...")
			session.Close()
			return nil
		},
	})

	// Initialize database middleware and shutdown routine
	diBuilder.Add(di.Def{
		Name: static.DiDatabase,
		Build: func(ctn di.Container) (interface{}, error) {
			return inits.InitDatabase(), nil
		},
		Close: func(obj interface{}) error {
			database := obj.(*mongo.Database)
			logrus.Info("Shutting down database connection...")
			database.Client().Disconnect(context.TODO())
			return nil
		},
	})

	// Initialize command handler
	diBuilder.Add(di.Def{
		Name: static.DiCommandHandler,
		Build: func(ctn di.Container) (interface{}, error) {
			return inits.InitCommandHandler(ctn)
		},
		Close: func(obj interface{}) error {
			logrus.Info("Unegister commands ...")
			return obj.(*ken.Ken).Unregister()
		},
	})

	// Build dependency injection container
	ctn := diBuilder.Build()

	// Setting log level from config
	logrus.SetLevel(logrus.Level(4))
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})

	// Initialize discord session and event
	// handlers
	inits.InitDiscordBotSession(ctn)

	// Block main go routine until one of the following
	// specified exit syscalls occure.
	logrus.Info("Bot started. Stop with CTRL-C...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Tear down dependency instances
	ctn.DeleteWithSubContainers()
}
