package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

//go:embed embed/cmd_info.md
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
	return "sp.etc.info"
}

func (c *CmdInfo) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdInfo) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdInfo) Exec(ctx shireikan.Context) error {
	st := ctx.GetObject(static.DiState).(*dgrs.State)
	self, err := st.SelfUser()
	if err != nil {
		return err
	}

	invLink := "https://discord.com/oauth2/authorize?client_id=901356147720749096&permissions=2080374975&scope=bot+applications.commands+identify+guilds"

	emb := &discordgo.MessageEmbed{
		Color: 0x5c32a8,
		Title: "Info",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: self.AvatarURL(""),
		},
		Description: infoMsg,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Repository",
				Value: "[github.com/archroid/ElProfessorBot](https://github.com/archroid/ElProfessorBot)",
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
				Name:  "Bug Hunters",
				Value: "Much :heart: to all the people who helped me with the bug hunting.",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("© %s archroid", time.Now().Format("2006")),
		},
	}
	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, emb)
	return err
}
