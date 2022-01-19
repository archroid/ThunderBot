package commands

import (
	"archroid/ElProfessorBot/static"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/zekroTJA/shireikan"
)

type CmdStop struct {
}

func (c *CmdStop) GetInvokes() []string {
	return []string{"stop", "s"}
}

func (c *CmdStop) GetDescription() string {
	return "stop the playing music"
}
func (c *CmdStop) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdStop) GetHelp() string {
	return "`stop`"
}

func (c *CmdStop) GetDomainName() string {
	return ""
}

func (c *CmdStop) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdStop) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdStop) Exec(ctx shireikan.Context) error {

	// session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	err := link.Player(ctx.GetGuild().ID).Stop()
	if err != nil {
		return err
	}

	return nil

}
