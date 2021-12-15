package music

import (
	"io"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
)

func Play(videoId string, voiceConnection *discordgo.VoiceConnection) (streamingSession *dca.StreamingSession, err error) {

	client := youtube.Client{}

	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = "lowdelay"

	video, err := client.GetVideo(videoId)
	if err != nil {
		return nil, err
	}

	formats := video.Formats.Type("audio").WithAudioChannels()

	downloadURL, err := client.GetStreamURL(video, &formats[0])
	if err != nil {
		return nil, err
	}

	encodingSession, err := dca.EncodeFile(downloadURL, options)
	if err != nil {
		return nil, err
	}
	defer encodingSession.Cleanup()

	done := make(chan error)
	streamingSession = dca.NewStream(encodingSession, voiceConnection, done)
	err = <-done
	if err != nil && err != io.EOF {
		return nil, err
	}
	return streamingSession, nil

}
