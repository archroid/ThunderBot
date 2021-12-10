package bot

import (
	"fmt"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	session *discordgo.Session
	db      *mongo.Database
)

func Start(token string, database *mongo.Database) {

	db = database

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Identify.Intents |= discordgo.IntentsGuildMembers
	session.Identify.Intents |= discordgo.IntentsDirectMessageReactions
	session.AddHandler(guildMemberAdd)
	session.AddHandler(ready)
	// session.AddHandler(command)

	session.AddHandler(message)

	defer session.Close()

	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}
	addCommands(session, commands)

	log.Printf("Bot is online: %v guilds \n", len(session.State.Guilds))

	// Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")

	// deleteAllCommands(session)

}
