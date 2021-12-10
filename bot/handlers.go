package bot

import (
	"archroid/ElProfessorBot/structs"
	embed "archroid/ElProfessorBot/utils"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func guildMemberAdd(session *discordgo.Session, member *discordgo.GuildMemberAdd) {
	var welcomeMessage structs.WelcomeMessage
	var autoRole structs.Role
	guildID := member.GuildID

	if db == nil {
		log.Println("db is nill")
	}

	filter := bson.M{"guildid": guildID}

	err := db.Collection("welcome").FindOne(context.TODO(), filter).Decode(&welcomeMessage)
	if err != nil {
		log.Println(err)
	} else {
		guild, _ := session.State.Guild(guildID)

		embed := embed.NewEmbed().
			SetColor(0x372168).
			SetThumbnail(guild.Icon).
			SetTitle("👋Welcome!").
			SetImage(member.User.AvatarURL("24")).
			SetDescription(fmt.Sprintf("Welcome to %v , %v \n %v", "this server", member.User.Username, welcomeMessage.WelcomeMessage)).
			MessageEmbed
		session.ChannelMessageSendEmbed(welcomeMessage.WelcomeChannelId, embed)
	}

	err = db.Collection("auto-role").FindOne(context.TODO(), filter).Decode(&autoRole)
	if err != nil {
		log.Printf("auto-role: %v", err)
	} else {
		err := session.GuildMemberRoleAdd(guildID, member.User.ID, autoRole.RoleID)
		if err != nil {
			log.Printf("auto-role: %v", err)
		}
	}

}

func ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(5, "/help")
}

func command(session *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(session, i)
	}
}

func message(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "&"):
		ping := bot.HeartbeatLatency().Truncate(60).Round(time.Millisecond)
		if message.Content == "&ping" {

			embed := embed.NewEmbed().
				SetColor(0xff0000).
				SetTitle("🏓").
				SetDescription(`Pong: **` + ping.String() + `** `).
				MessageEmbed
			bot.ChannelMessageSendEmbed(message.ChannelID, embed)
		}

		if message.Content == "&github" {
			embed := embed.NewEmbed().
				SetAuthor(message.Author.Username, message.Author.AvatarURL("1024")).
				SetThumbnail(message.Author.AvatarURL("1024")).
				SetTitle("My repository").
				SetDescription("You can find my repository by clicking [here](https://github.com/archroid).").
				SetColor(0x00ff00).MessageEmbed
			bot.ChannelMessageSendEmbed(message.ChannelID, embed)
		}
		if message.Content == "&botinfo" {
			guilds := len(bot.State.Guilds)
			embed := embed.NewEmbed().
				SetTitle("ElProfessor Bot").
				SetColor(0x372168).
				SetThumbnail("https://cdn.discordapp.com/avatars/901356147720749096/3107c752e9bc40bcb9dd0100bd53976b.png").
				SetDescription("Some informations about me :)").
				SetAuthor("Professor#9681", "https://cdn.discordapp.com/avatars/782162374890487810/32a321b1b588f2126aec41b833030590.png").
				AddField("GO version:", runtime.Version()).
				AddField("DiscordGO version:", discordgo.VERSION).
				AddField("Concurrent tasks:", strconv.Itoa(runtime.NumGoroutine())).
				AddField("📡Latency:", ping.String()).
				AddField("Author:", "Made with ❤️ by Professor#9681").
				AddField("Invitation Link:", "https://b2n.ir/n97207").
				AddField("Total guilds:", strconv.Itoa(guilds)).MessageEmbed
			bot.ChannelMessageSendEmbed(message.ChannelID, embed)

		}

		if message.Content == "&clear" {

			st, err := bot.ChannelMessages(message.ChannelID, 99, message.Reference().MessageID, "", "")
			if err != nil {
				log.Panicln(err)
				return
			}

			var messageIds []string
			for _, strings := range st {
				messageIds = append(messageIds, strings.Reference().MessageID)
			}

			messageIds = append(messageIds, message.Reference().MessageID)

			log.Printf("msgid %v", 0)
			log.Printf("%v messages deleted \n", len(messageIds))

			//Delete messages
			bot.ChannelMessagesBulkDelete(message.ChannelID, messageIds)

			//Say the user about deleted messagess
			embed := embed.NewEmbed().
				SetTitle(fmt.Sprintf("%v messages has been deleted!", len(messageIds))).
				SetColor(0xff0000).
				MessageEmbed
			bot.ChannelMessageSendEmbed(message.ChannelID, embed)

			//Delete the message itself
			embedMessage, err := bot.ChannelMessages(message.ChannelID, 1, "", "", "")
			if err != nil {
				log.Panicln(err)
				return
			}
			println(len(embedMessage))
			embedMessageString := embedMessage[0].Reference().MessageID

			//wait 3 seconds
			time.Sleep(time.Second * 2)

			bot.ChannelMessageDelete(message.ChannelID, embedMessageString)

		}
	}
}
