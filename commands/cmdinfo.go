package commands

import (
	"archroid/ElProfessorBot/static"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

var infoMsg string

type CmdInfo struct {
}

func (c *CmdInfo) GetInvokes() []string {
	return []string{"info", "information", "description", "credits", "version", "invite"}
}

func (c *CmdInfo) GetDescription() string {
	return "Display some information about this bot."
}

func (c *CmdInfo) GetHelp() string {
	return "`info`"
}

func (c *CmdInfo) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdInfo) GetDomainName() string {
	return ""
}

func (c *CmdInfo) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdInfo) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdInfo) Exec(ctx shireikan.Context) error {

	session := ctx.GetSession()

	invLink := static.DiscordInviteLink

	emb := &discordgo.MessageEmbed{
		Color: static.ColorEmbedDefault,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: session.State.User.AvatarURL(""),
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    session.State.User.Username,
			IconURL: session.State.User.AvatarURL(""),
		},
		Description: infoMsg,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🏅 Version",
				Value:  "0.0.1 Alpha",
				Inline: true,
			},
			{
				Name:   "📢 Servers",
				Value:  strconv.Itoa(len(session.State.Guilds)),
				Inline: true,
			},

			{
				Name:   "💻 Created by",
				Value:  "@Thuйdёr#0477",
				Inline: true,
			},
			{
				Name:  "🚥 Licence",
				Value: "Covered by the [GNU General Public License](https://github.com/archroid/ThunderBot/blob/main/LICENSE).",
			},
			{
				Name:  "📑 Repository",
				Value: "[github.com/archroid/ElProfessorBot](https://github.com/archroid/ThunderBot)",
			},

			{
				Name:  "🌀 Website",
				Value: "[elprofessorbot.archroid.xyz](https://ThunderBot.archroid.xyz)",
			},
			{
				Name: "📡 Invite",
				Value: fmt.Sprintf("[Invite Link](%s)",
					invLink),
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("© 2021-%s archroid", time.Now().Format("2006")),
		},
	}
	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, emb)
	return err
}
