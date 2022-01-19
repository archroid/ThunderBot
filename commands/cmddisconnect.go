package commands

import (
	"archroid/ElProfessorBot/static"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/zekroTJA/shireikan"
)

type CmdDisconnect struct {
}

func (c *CmdDisconnect) GetInvokes() []string {
	return []string{"disconnect"}
}

func (c *CmdDisconnect) GetDescription() string {
	return "disconnect from the current voice channel"
}
func (c *CmdDisconnect) GetGroup() string {
	return shireikan.GroupGeneral
}

func (c *CmdDisconnect) GetHelp() string {
	return "`disconnect`"
}

func (c *CmdDisconnect) GetDomainName() string {
	return ""
}

func (c *CmdDisconnect) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdDisconnect) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdDisconnect) Exec(ctx shireikan.Context) error {

	// session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	err := link.Player(ctx.GetGuild().ID).Destroy()
	if err != nil {
		return err
	}

	return nil

}
