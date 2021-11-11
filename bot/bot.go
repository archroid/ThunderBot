package bot

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"test/Discord-Template/config"
	embed "test/Discord-Template/utils"
	"time"

	"github.com/bwmarrin/discordgo"
)

var botID string
var client *discordgo.Session

func Start() {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	session.AddHandler(message)
	session.AddHandler(ready)

	defer session.Close()
	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Bot is online")

	scall := make(chan os.Signal, 1)
	signal.Notify(scall, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-scall
}

func ready(bot *discordgo.Session, event *discordgo.Ready) {
	guildsSize := len(bot.State.Guilds)
	bot.UpdateGameStatus(0, strconv.Itoa(guildsSize)+" guilds!")
}

func message(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	switch {
	case strings.HasPrefix(message.Content, config.BotPrefix):
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
	}
}
