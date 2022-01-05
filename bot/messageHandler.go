package bot

import (
	embed "archroid/ElProfessorBot/utils"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func message(s *discordgo.Session, message *discordgo.MessageCreate) {

	prefix := "-"

	//Check if the message is send by the bot
	if message.Author.Bot {
		return
	}

	//Check if the message has Prefix
	if !strings.HasPrefix(message.Content, prefix) {
		return
	}

	switch message.Content {

	// ping
	case prefix + "ping":
		ping := s.HeartbeatLatency().Truncate(60).Round(time.Millisecond)
		embed := embed.NewEmbed().
			SetColor(0xff0000).
			SetTitle("pong!🏓").
			SetDescription(`⌛** Time: **` + ping.String()).
			MessageEmbed
		s.ChannelMessageSendEmbed(message.ChannelID, embed)

	//help
	case (prefix + "help"), (prefix + "commands"):
		s.ChannelMessageSend(message.ChannelID, "https://elprofessorbot.archroid.xyz/commands")

	//clear
	case (prefix + "clear"):
		var deleteNum int64 = 100

		st, err := s.ChannelMessages(message.ChannelID, int(deleteNum), "", "", "")
		if err != nil {
			log.Panicln(err)
			return
		}
		var messageIds []string
		for _, strings := range st {
			messageTimestamp, _ := strings.Timestamp.Parse()
			messageTimestampUnix := messageTimestamp.Unix()

			twoWeeksTimestampUnix := time.Now().AddDate(0, 0, -14).Unix()

			if messageTimestampUnix >= twoWeeksTimestampUnix {
				messageIds = append(messageIds, strings.Reference().MessageID)
			}
		}

		returnText := fmt.Sprintf("%v message(s) has been deleted!", len(messageIds))

		err = s.ChannelMessagesBulkDelete(message.ChannelID, messageIds)

		if err != nil {
			returnText = "You can only delete messages that are under 14 days old."
		}

		embed := embed.NewEmbed().
			SetTitle(returnText).
			SetColor(0xff0000).
			MessageEmbed
		sentMessage, _ := s.ChannelMessageSendEmbed(message.ChannelID, embed)

		time.Sleep(time.Second * 2)
		s.ChannelMessageDelete(message.ChannelID, sentMessage.ID)
	}

}
