package youtubemusic

import (
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"

	"github.com/jonas747/dca"
)

func PlayMusic(videoID string, vc *discordgo.VoiceConnection) (err error) {
	encodeSession, err := dca.EncodeFile(PathToAudio(videoID), dca.StdEncodeOptions)
	if err != nil {
		log.Print(err)
		return err
	}
	defer encodeSession.Cleanup()

	decoder := dca.NewDecoder(encodeSession)
	println(decoder.FrameDuration())
	// abortChannel := ActiveVoiceChannels[guildID].GetAbortChannel()

	// for {
	// 	frame, err := decoder.OpusFrame()
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Print(err)
	// 			return err
	// 		}
	// 		break
	// 	}

	// 	select {
	// 	case vc.OpusSend <- frame:
	// 	case <-time.After(time.Second):
	// 		// We haven't been able to send a frame in a second, assume the connection is borked
	// 		return err
	// 	}

	// }

	// vc.Speaking(true)

	return
}
