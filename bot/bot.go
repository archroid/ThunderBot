package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	session *discordgo.Session
	db      *mongo.Database
)

func Start(token string, database *mongo.Database) {

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.AddHandler(guildMemberAdd)
	session.AddHandler(ready)
	session.AddHandler(command)
	session.Identify.Intents |= discordgo.IntentsGuildMembers
	session.Identify.Intents |= discordgo.IntentsDirectMessageReactions
	// session.AddHandler(message)

	defer session.Close()
	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Bot is online: %v guilds \n", len(session.State.Guilds))

	addCommands(session, commands)

	db = database

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
	deleteAllCommands(session)

}
