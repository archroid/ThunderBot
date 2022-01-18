package main

import (
	"archroid/ElProfessorBot/inits"
	"archroid/ElProfessorBot/static"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/DisgoOrg/disgolink/dgolink"

	"github.com/DisgoOrg/disgolink/lavalink"
	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/zekroTJA/shireikan"
	"github.com/zekrotja/ken"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	// Initialize dependency injection builder
	diBuilder, _ := di.NewBuilder()

	// Initialize discord bot session and shutdown routine
	diBuilder.Add(di.Def{
		Name: static.DiContainer,
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn, nil
		},
	})

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

	// Initialize dgolink
	diBuilder.Add(di.Def{
		Name: static.DiDgoLink,
		Build: func(ctn di.Container) (interface{}, error) {

			link := dgolink.New(ctn.Get(static.DiDiscordSession).(*discordgo.Session))

			link.AddNode(lavalink.NodeConfig{
				Name:     "test",
				Host:     "localhost",
				Port:     "2333",
				Password: "1274",
			})

			return link, nil
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

	// Initialize legacy command handler
	diBuilder.Add(di.Def{
		Name: static.DiLegacyCommandHandler,
		Build: func(ctn di.Container) (interface{}, error) {
			return inits.InitLegacyCommandHandler(ctn), nil
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

	// ctn.Get(static.DiCommandHandler)

	// Initialize discord session and event
	// handlers
	inits.InitDiscordBotSession(ctn)

	// This is currently the really hacky workaround
	// to bypass the di.Container when trying to get
	// the Command legacyHandler instance inside a command
	// context, because the legacyHandler can not resolve
	// itself on build, so it is bypassed here using
	// shireikans object map. Maybe I find a better
	// solution for that at some time.
	legacyHandler := ctn.Get(static.DiLegacyCommandHandler).(shireikan.Handler)
	legacyHandler.SetObject(static.DiLegacyCommandHandler, legacyHandler)

	// Block main go routine until one of the following
	// specified exit syscalls occure.
	logrus.Info("Bot started. Stop with CTRL-C...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Tear down dependency instances
	ctn.DeleteWithSubContainers()
}
