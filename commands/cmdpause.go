package commands

import (
	"archroid/ElProfessorBot/static"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/zekroTJA/shireikan"
)

type CmdPause struct {
}

func (c *CmdPause) GetInvokes() []string {
	return []string{"pause"}
}

func (c *CmdPause) GetDescription() string {
	return "get the ping of the bot"
}
func (c *CmdPause) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdPause) GetHelp() string {
	return "`pause`"
}

func (c *CmdPause) GetDomainName() string {
	return ""
}

func (c *CmdPause) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdPause) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdPause) Exec(ctx shireikan.Context) error {

	// session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	link.Player(ctx.GetGuild().ID).Pause(true)
	return nil

}
