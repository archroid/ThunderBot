package commands

import (
	"archroid/ElProfessorBot/static"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/zekroTJA/shireikan"
)

type CmdResume struct {
}

func (c *CmdResume) GetInvokes() []string {
	return []string{"resume", "r"}
}

func (c *CmdResume) GetDescription() string {
	return "resume the currently playing music"
}
func (c *CmdResume) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdResume) GetHelp() string {
	return "`resume`"
}

func (c *CmdResume) GetDomainName() string {
	return ""
}

func (c *CmdResume) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdResume) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdResume) Exec(ctx shireikan.Context) error {

	// session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	link.Player(ctx.GetGuild().ID).Pause(false)

	return nil

}
