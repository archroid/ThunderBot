package playservice

import (
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
)

func PlayYoutube(videoId string, voiceConnection *discordgo.VoiceConnection) (err error) {
	client := youtube.Client{}

	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 64
	// options.Application = "lowdelay"

	video, err := client.GetVideo(videoId)
	if err != nil {
		return
	}
	formats := video.Formats.Type("audio").WithAudioChannels()

	downloadURL, err := client.GetStreamURL(video, &formats[0])
	if err != nil {
		return
	}

	encodingSession, err := dca.EncodeFile(downloadURL, options)
	if err != nil {
		return
	}

	done := make(chan error)
	stream := dca.NewStream(encodingSession, voiceConnection, done)

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case err = <-done:
			if err != nil && err != io.EOF {
				logrus.WithError(err).Error("Error playing stream")
			}

			// Clean up incase something happened and ffmpeg is still running
			encodingSession.Truncate()
			return
		case <-ticker.C:
			stats := encodingSession.Stats()
			playbackPosition := stream.PlaybackPosition()

			fmt.Printf("\nPlayback: %10s, Transcode Stats: Time: %5s, Size: %5dkB, Bitrate: %6.2fkB, Speed: %5.1fx\r", playbackPosition, stats.Duration.String(), stats.Size, stats.Bitrate, stats.Speed)
		}
	}
}
