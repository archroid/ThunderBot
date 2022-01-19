package commands

import (
	"archroid/ElProfessorBot/static"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/zekroTJA/shireikan"
)

type CmdJoin struct {
}

func (c *CmdJoin) GetInvokes() []string {
	return []string{"pause"}
}

func (c *CmdJoin) GetDescription() string {
	return "get the ping of the bot"
}
func (c *CmdJoin) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdJoin) GetHelp() string {
	return "`pause`"
}

func (c *CmdJoin) GetDomainName() string {
	return ""
}

func (c *CmdJoin) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdJoin) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdJoin) Exec(ctx shireikan.Context) error {

	// session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	link.Player(ctx.GetGuild().ID).Pause(true)
	return nil

}
