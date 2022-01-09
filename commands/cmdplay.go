package commands

import (
	"archroid/ElProfessorBot/searchservice"

	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/shireikan"
)

type CmdPlay struct {
}

func (c *CmdPlay) GetInvokes() []string {
	return []string{"play", "p"}
}

func (c *CmdPlay) GetDescription() string {
	return "Play a music from youtube"
}

func (c *CmdPlay) GetHelp() string {
	return "`clear <query>` - Play by youtube url or search"
}

func (c *CmdPlay) GetGroup() string {
	return shireikan.GroupModeration
}

func (c *CmdPlay) GetDomainName() string {
	return ""
}

func (c *CmdPlay) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdPlay) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdPlay) Exec(ctx shireikan.Context) error {

	if len(ctx.GetArgs()) == 0 {
		return nil
	}

	query := ctx.GetArgs()[0]

	videoId, err := searchservice.GetVideoID(query, ctx.GetObject("container").(di.Container))
	if err != nil {
		return err
	}

	ctx.GetSession().ChannelMessageSend(ctx.GetChannel().ID, "https://www.youtube.com/watch?v="+videoId)
	return nil

}
