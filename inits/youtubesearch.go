package inits

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func InitYoutubeSearch() *youtube.SearchService {
	context := context.Background()
	youtubeService, err := youtube.NewService(context, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Sucessfully initialised Youtube Search Service.")
	return youtubeService.Search
}
