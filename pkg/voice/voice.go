package voice

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func JoinVoice(s *discordgo.Session, guild *discordgo.Guild, user *discordgo.User) (vc *discordgo.VoiceConnection, err error) {

	var channel *discordgo.Channel

	if GetCurrentVoiceChannel(user, s, guild) == nil {
		return nil, fmt.Errorf("you are not connected to a voice channel")
	} else {
		channel = GetCurrentVoiceChannel(user, s, guild)
		VoiceConnection, err := s.ChannelVoiceJoin(guild.ID, channel.ID, false, true)
		if err != nil {
			return nil, err
		}
		return VoiceConnection, nil
	}
}

func DisconnectVoice(vc *discordgo.VoiceConnection) (err error) {
	err = vc.Disconnect()
	if err != nil {

		return err
	}
	return nil
}

func GetCurrentVoiceChannel(user *discordgo.User, session *discordgo.Session, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == user.ID {
			channel, _ := session.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}

