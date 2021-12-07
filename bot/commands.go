package bot

import (
	"archroid/ElProfessorBot/structs"
	embed "archroid/ElProfessorBot/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	guildCommand string = "801840788022624296"

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

		{

			Name:        "welcome",
			Description: "Enable welcome message on user join",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "enabled",
					Description: "Enable or disable the welcoming system.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "welcome-channel",
					Description: "The text channel you want your welcome messages send to them.",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "welcome-message",
					Description: "The text you want to send as welcome message.",
					Required:    false,
				},
			},
		},

		{
			Name:        "auto-role",
			Description: "give a special role to anyone that joins the crew!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "enabled",
					Description: "Enable or disable auto-roling system.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role you want to set.",
					Required:    false,
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
		"welcome": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			if i.ApplicationCommandData().Options[0].BoolValue() {

				if len(i.ApplicationCommandData().Options) == 3 {
					welcomeChannelId := i.ApplicationCommandData().Options[1].ChannelValue(session).ID
					welcomeMessage := i.ApplicationCommandData().Options[2].StringValue()

					guildId := i.GuildID

					//remove previous settings from database
					filter := bson.M{"guildid": guildId}

					_, err := db.Collection("welcome").DeleteOne(context.TODO(), filter)
					if err != nil {
						log.Println(err)
					}

					insertWelcome := structs.WelcomeMessage{welcomeChannelId, welcomeMessage, guildId}
					_, err = db.Collection("welcome").InsertOne(context.TODO(), insertWelcome)
					if err != nil {
						embed := embed.NewEmbed().
							SetColor(0xff0000).
							SetTitle("🔴Error!").
							SetDescription(`Error setting welcome message`).
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
					} else {
						embed := embed.NewEmbed().
							SetColor(0x00ff00).
							SetTitle("✅Done!").
							SetDescription(`Your welcome message settings saved!`).
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
					}
				} else {
					embed := embed.NewEmbed().
						SetColor(0xff0000).
						SetTitle("🔴Error!").
						SetDescription(`You should fill all fields!`).
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
				}

			} else {
				guildId := i.GuildID

				filter := bson.M{"guildid": guildId}

				_, err := db.Collection("welcome").DeleteOne(context.TODO(), filter)
				if err != nil {
					log.Println(err)
				}

				embed := embed.NewEmbed().
					SetColor(0x00ff00).
					SetTitle("🟢Done!").
					SetDescription(`I won't send welcome messages to this server anymore!`).
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
			}
		},
		"auto-role": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.ApplicationCommandData().Options[0].BoolValue() {
				if len(i.ApplicationCommandData().Options) == 2 {
					guildId := i.GuildID
					roleID := i.ApplicationCommandData().Options[1].RoleValue(session, guildId).ID

					//remove previous settings from database
					filter := bson.M{"guildid": guildId}

					_, err := db.Collection("auto-role").DeleteOne(context.TODO(), filter)
					if err != nil {
						log.Println(err)
					}
					insertRole := structs.Role{roleID, guildId}
					_, err = db.Collection("auto-role").InsertOne(context.TODO(), insertRole)
					if err != nil {
						embed := embed.NewEmbed().
							SetColor(0xff0000).
							SetTitle("🔴Error!").
							SetDescription(`Error setting welcome message`).
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
					} else {
						embed := embed.NewEmbed().
							SetColor(0x00ff00).
							SetTitle("✅Done!").
							SetDescription(`auto-role settings saved!`).
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
					}
				} else {
					embed := embed.NewEmbed().
						SetColor(0xff0000).
						SetTitle("🔴Error!").
						SetDescription(`You should fill all fields!`).
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
				}
			} else {
				guildId := i.GuildID

				filter := bson.M{"guildid": guildId}

				_, err := db.Collection("auto-role").DeleteOne(context.TODO(), filter)
				if err != nil {
					log.Println(err)
				}

				embed := embed.NewEmbed().
					SetColor(0x00ff00).
					SetTitle("🟢Done!").
					SetDescription(`I won't give roles to members when they join on this server anymore!`).
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
			}
		},
	}
)

func addCommands(session *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildCommand, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

}

func deleteAllCommands(session *discordgo.Session) {
	commands, _ := session.ApplicationCommands(session.State.User.ID, guildCommand)
	for _, v := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, guildCommand, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
