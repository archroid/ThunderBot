package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

func Start(token string) {

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.AddHandler(guildMemberAdd)
	session.AddHandler(ready)
	session.AddHandler(command)
	// session.AddHandler(message)

	defer session.Close()
	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Bot is online: %v guilds \n", len(session.State.Guilds))

	addCommands(session, commands)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
	deleteAllCommands(session)

}


