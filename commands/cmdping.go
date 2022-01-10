package commands

import (
	"archroid/ElProfessorBot/static"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

type CmdPing struct {
}

func (c *CmdPing) GetInvokes() []string {
	return []string{"ping"}
}

func (c *CmdPing) GetDescription() string {
	return "get the ping of the bot"
}
func (c *CmdPing) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdPing) GetHelp() string {
	return "`ping`"
}

func (c *CmdPing) GetDomainName() string {
	return ""
}

func (c *CmdPing) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdPing) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdPing) Exec(ctx shireikan.Context) error {

	s := ctx.GetSession()
	ping := s.HeartbeatLatency().Truncate(60).Round(time.Millisecond)

	emb := &discordgo.MessageEmbed{
		Color:       static.ColorEmbedGreen,
		Title:       "💓 Pong",
		Description: `📡** Letancy: **` + ping.String(),
	}
	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, emb)
	return err
}
