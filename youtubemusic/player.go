package youtubemusic

import (
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"

	"github.com/jonas747/dca"
)

func PlayMusic(videoID string, vc *discordgo.VoiceConnection) {
	encodeSession, err := dca.EncodeFile(PathToAudio(videoID), dca.StdEncodeOptions)
	if err != nil {
		log.Print(err)
		return 
	}
	defer encodeSession.Cleanup()

	decoder := dca.NewDecoder(encodeSession)
	println(decoder.FrameDuration())

}
