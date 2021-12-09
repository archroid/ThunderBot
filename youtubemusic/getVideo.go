package youtubemusic

import "github.com/kkdai/youtube/v2"

func getVideo() {
	videoID := "go7IbOfEYh4"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	println(video.Title)
}
