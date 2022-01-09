package commands

import (
	"archroid/ElProfessorBot/static"
	"fmt"
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
	return "archroid.xyz"
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
		Title: "Info",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: session.State.User.AvatarURL(""),
		},
		Description: infoMsg,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Repository",
				Value: "[github.com/zekrotja/shinpuru](https://github.com/zekrotja/shinpuru)",
			},
			{
				Name:  "Licence",
				Value: "Covered by the [GNU General Public License](https://github.com/archroid/ElProfessorBot/blob/main/LICENSE).",
			},
			{
				Name: "Invite",
				Value: fmt.Sprintf("[Invite Link](%s).\n```\n%s\n```",
					invLink, invLink),
			},
			{
				Name:  "Development state",
				Value: "Testing.",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("© 2021-%s archroid", time.Now().Format("2006")),
		},
	}
	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, emb)
	return err
}
