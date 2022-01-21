package music

import (
	"archroid/ElProfessorBot/static"
	"archroid/ElProfessorBot/utils"
	"fmt"
	"time"

	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

func Play(ctx shireikan.Context) {

	s := ctx.GetSession()
	guild := ctx.GetGuild()
	plManager := ctx.GetObject(static.DiPlaylistManager).(PlaylistManager)
	link := ctx.GetObject(static.DiDgoLink).(*dgolink.Link)

	channel, err := getCurrentVoiceChannel(ctx.GetUser().ID, s, guild)
	if err != nil {
		utils.SendEmbedError(s, ctx.GetChannel().ID,
			"Couldn't find you in any voice channel. Are you connected?").
			DeleteAfter(5 * time.Second).Error()
		return
	}

	err = s.ChannelVoiceJoinManual(guild.ID, channel.ID, false, false)
	if err != nil {
		utils.SendEmbedError(s, ctx.GetChannel().ID,
			"Couldn't connect to the voice channel.").
			DeleteAfter(5 * time.Second).Error()
		return
	}

	for tracks := plManager.GetPlaylist(guild.ID).Tracks; len(tracks) > 0; {
		link.Player(guild.ID).Play(tracks[0])
		utils.SendEmbed(s,
			ctx.GetChannel().ID,
			fmt.Sprint("Playing %s", link.Player(guild.ID).Track().Info().Title()),
			"🌀Playing",
			static.ColorEmbedDefault)
			
		plManager.RemoveLastTrack(guild.ID)
		tracks = plManager.GetPlaylist(guild.ID).Tracks
	}

}

func getCurrentVoiceChannel(userId string, session *discordgo.Session, guild *discordgo.Guild) (chnl *discordgo.Channel, err error) {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == userId {
			channel, _ := session.Channel(vs.ChannelID)
			return channel, nil
		}
	}
	return nil, fmt.Errorf("you are no connected to any voice channel")
}
