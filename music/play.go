package music

import (
	"github.com/DisgoOrg/disgolink/dgolink"
	"github.com/DisgoOrg/disgolink/lavalink"
	"github.com/bwmarrin/discordgo"
)

func Play(s *discordgo.Session, link *dgolink.Link, guild *discordgo.Guild, userId string, track lavalink.Track) {

	channelID := getCurrentVoiceChannel(userId, s, guild).ID

	if err := s.ChannelVoiceJoinManual(guild.ID, channelID, false, false); err != nil {
		_, _ = s.ChannelMessageSend(channelID, "Error while joining voice channel: "+err.Error())
		return
	}
	if err := link.Player(guild.ID).Play(track); err != nil {
		_, err = s.ChannelMessageSend(channelID, "Error while playing track: "+err.Error())
		return
	}
	_, _ = s.ChannelMessageSend(channelID, "Playing: "+track.Info().Title())

}

func getCurrentVoiceChannel(userId string, session *discordgo.Session, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == userId {
			channel, _ := session.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}
