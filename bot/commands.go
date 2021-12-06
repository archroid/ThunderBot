package bot

import (
	embed "archroid/ElProfessorBot/utils"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	guildID string = "801840788022624296"

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
					Required:    false,
				},
			},
		},

		// {

		// 	Name: "",
		// },
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

			var deleteNum int64 = 100

			if len(i.ApplicationCommandData().Options) == 1 {
				deleteNum = i.ApplicationCommandData().Options[0].IntValue()
			}

			st, err := s.ChannelMessages(i.ChannelID, int(deleteNum), "", "", "")
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

			deletedMSG := fmt.Sprintf("%v message(s) has been deleted!", len(messageIds))

			err = s.ChannelMessagesBulkDelete(i.ChannelID, messageIds)

			if err != nil {
				deletedMSG = "You can only bulk delete messages that are under 14 days old."
			}

			// log.Println(deletedMSG)

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

func addCommands(session *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

}

func deleteAllCommands(session *discordgo.Session) {
	commands, _ := session.ApplicationCommands(session.State.User.ID, guildID)
	for _, v := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, guildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
