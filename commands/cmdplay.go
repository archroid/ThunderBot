package commands

import (
	"archroid/ElProfessorBot/static"
	"strings"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/sirupsen/logrus"

	"github.com/DisgoOrg/disgolink/lavalink"
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
	return "`play <query>` - Play by youtube url or search"
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

	query := strings.Join(ctx.GetArgs()[:], " ")

	if !static.UrlPattern.MatchString(query) {
		query = "ytsearch:" + query
	}

	session := ctx.GetSession()

	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	link.BestRestClient().LoadItemHandler(query, lavalink.NewResultHandler(
		func(track lavalink.Track) {
			// music.Play(session, link, ctx.GetGuild(), ctx.GetUser().ID, track)
		},
		func(playlist lavalink.Playlist) {
			// music.Play(session, link, ctx.GetGuild(), ctx.GetUser().ID, playlist.Tracks[0])
		},
		func(tracks []lavalink.Track) {
			// music.Play(session, link, ctx.GetGuild(), ctx.GetUser().ID, tracks[0])
		},
		func() {
			_, err := session.ChannelMessageSend(ctx.GetChannel().ID, "No matches found for: "+query)
			if err != nil {
				logrus.Info(err)
			}
		},
		func(ex lavalink.Exception) {
			_, err := session.ChannelMessageSend(ctx.GetChannel().ID, "Error while loading track: "+ex.Message)
			if err != nil {
				logrus.Info(err)
			}
		},
	))

	return nil

}
