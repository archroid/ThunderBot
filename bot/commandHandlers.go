package bot

import (
	"archroid/ElProfessorBot/music"
	"archroid/ElProfessorBot/poll"
	"archroid/ElProfessorBot/structs"
	embed "archroid/ElProfessorBot/utils"
	"archroid/ElProfessorBot/youtubemusic"
	"context"
	"fmt"

	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

var vc *discordgo.VoiceConnection

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
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
			deletedMSG = "You can only delete messages that are under 14 days old."
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
	"set-welcome": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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

				// insertWelcome := structs.WelcomeMessage{welcomeChannelId, welcomeMessage, guildId}
				insertWelcome := structs.WelcomeMessage{WelcomeChannelId: welcomeChannelId, WelcomeMessage: welcomeMessage, GuildId: guildId}
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
				// insertRole := structs.Role{roleID, guildId}
				insertRole := structs.Role{RoleID: roleID, GuildID: guildId}
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
	"roll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		RandomIntegerwithinRange := rand.Intn(int(i.ApplicationCommandData().Options[0].IntValue()))
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprint(RandomIntegerwithinRange),
			},
		})
	},

	"set-rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		rules := i.ApplicationCommandData().Options[0].StringValue()

		filter := bson.M{"guildid": i.GuildID}

		_, err := db.Collection("rules").DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println(err)
		}

		// insertRules := structs.Rules{rules, i.GuildID}
		insertRules := structs.Rules{Rules: rules, GuildID: i.GuildID}

		_, err = db.Collection("rules").InsertOne(context.TODO(), insertRules)
		if err != nil {
			embed := embed.NewEmbed().
				SetColor(0xff0000).
				SetTitle("🔴Error!").
				SetDescription(`Error saving rules.`).
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
				SetDescription(`Rules saved!`).
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

	"rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var rules structs.Rules

		filter := bson.M{"guildid": i.GuildID}

		err := db.Collection("rules").FindOne(context.TODO(), filter).Decode(&rules)
		if err != nil {
			embed := embed.NewEmbed().
				SetColor(0xff0000).
				SetTitle("🔴Error!").
				SetDescription(`This server has no rules`).
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: rules.Rules,
				},
			})
		}
	},

	"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://elprofessorbot.archroid.xyz",
			},
		})

	},

	"join": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		joinVoice(s, i)
	},

	"disconnect": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		vc.Disconnect()

		embed := embed.NewEmbed().
			SetColor(0x00ff00).
			SetTitle("✅Done!").
			SetDescription(`Disconnected`).
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

	"play": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		videoId, err := music.GetVideoID(i.ApplicationCommandData().Options[0].StringValue())
		if err != nil {
			log.Println(err)
		}

		err = music.GetVideoDownloadUrl(videoId, vc)
		if err != nil {
			log.Println(err)
		}
	},
	"stop": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	},
	"search": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		musicInput := i.ApplicationCommandData().Options[0].StringValue()

		videoID, err := youtubemusic.GetVideoID(musicInput)
		if err != nil {
			log.Println(err)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("https://www.youtube.com/watch?v=%v", videoID),
			},
		})
	},
	"poll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		poll.CreatePoll(s, i)
	},
}

func joinVoice(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, _ := s.State.Guild(i.GuildID)
	var channel *discordgo.Channel

	if getCurrentVoiceChannel(i.Member.User, s, guild) == nil {
		log.Println("You are not connected to a voice channel.")
		embed := embed.NewEmbed().
			SetColor(0xff0000).
			SetTitle("🔴Error!").
			SetDescription(`You are not connected to a voice channel.`).
			MessageEmbed

		embeds := []*discordgo.MessageEmbed{embed}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "",
				Embeds:  embeds,
			},
		})
		time.Sleep(time.Second * 3)
		s.InteractionResponseDelete(s.State.User.ID, i.Interaction)
	} else {
		channel = getCurrentVoiceChannel(i.Member.User, s, guild)
		VoiceConnection, err := s.ChannelVoiceJoin(i.GuildID, channel.ID, false, true)
		if err != nil {
			log.Println(err)
			embed := embed.NewEmbed().
				SetColor(0xff0000).
				SetTitle("🔴Error!").
				SetDescription(`Error joining the voice`).
				MessageEmbed

			embeds := []*discordgo.MessageEmbed{embed}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "",
					Embeds:  embeds,
				},
			})
			time.Sleep(time.Second * 3)
			s.InteractionResponseDelete(s.State.User.ID, i.Interaction)

		} else {
			vc = VoiceConnection
		}
	}
}

func getCurrentVoiceChannel(user *discordgo.User, session *discordgo.Session, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == user.ID {
			channel, _ := session.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}
