package music

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
)

func GetVideoDownloadUrl(videoId string, voiceConnection *discordgo.VoiceConnection) (err error) {

	client := youtube.Client{}
	// Change these accordingly
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = "lowdelay"

	video, err := client.GetVideo(videoId)
	if err != nil {
		return err
	}

	formats := video.Formats.Type("audio").WithAudioChannels()

	downloadURL, err := client.GetStreamURL(video, &formats[0])
	if err != nil {
		return err
	}
	log.Println(downloadURL)

	encodingSession, err := dca.EncodeFile(downloadURL, options)
	if err != nil {
		return err
	}
	defer encodingSession.Cleanup()

	done := make(chan error)
	dca.NewStream(encodingSession, voiceConnection, done)
	err = <-done
	if err != nil && err != io.EOF {
		return err
	}
	return nil

}
