package bot

import (
	embed "archroid/ElProfessorBot/utils"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "ping",

			Description: "Get the Bot's ping.",
		},

		{
			Name:        "clear",
			Description: "Removes latest messages.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number-of-messages",
					Description: "Number of messages to delete(max 100)",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			ping := s.HeartbeatLatency().Truncate(60).Round(time.Millisecond)
			embed := embed.NewEmbed().
				SetColor(0xff0000).
				SetTitle("pong!🏓").
				SetDescription(`⌛** Time: **` + ping.String()).
				MessageEmbed

			embeds := []*discordgo.MessageEmbed{embed}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "",
					Embeds:  embeds,
				},
			})
		},

		"clear": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			deleteNum := i.ApplicationCommandData().Options[0].IntValue()

			st, err := s.ChannelMessages(i.ChannelID, int(deleteNum), "", "", "")
			if err != nil {
				log.Panicln(err)
				return
			}
			var messageIds []string
			for _, strings := range st {

				// println(strings.Timestamp)
				// if strings.Timestamp >= discordgo.Timestamp() {
				// }
				messageIds = append(messageIds, strings.Reference().MessageID)
			}

			deletedMSG := fmt.Sprintf("%v messages has been deleted!", len(messageIds))

			err = s.ChannelMessagesBulkDelete(i.ChannelID, messageIds)

			if err != nil {
				deletedMSG = "You can only bulk delete messages that are under 14 days old."
			}

			log.Printf("%v messages deleted \n", deletedMSG)

			embed := embed.NewEmbed().
				SetTitle(deletedMSG).
				SetColor(0xff0000).
				MessageEmbed
			embeds := []*discordgo.MessageEmbed{embed}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "",
					Embeds:  embeds,
				},
			})

			time.Sleep(time.Second * 2)
			s.InteractionResponseDelete(s.State.User.ID, i.Interaction)
		},
	}
)
