package music

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

var searchService *youtube.SearchService

func initYtSearch() {
	context := context.Background()
	youtubeService, err := youtube.NewService(context, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	searchService = youtubeService.Search
	log.Info("Sucessfully initialised Youtube Search Service.")
}

func GetVideoID(query string) (youtubeID string, err error) {
	if searchService == nil {
		initYtSearch()
	}

	call := searchService.List([]string{"id, snippet"}).
		Type("video").
		Q(query).
		MaxResults(1)

	res, err := call.Do()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(res.Items) == 0 {
		err = fmt.Errorf("no results could be found for your query: %s", query)
		return
	}

	result := res.Items[0]
	youtubeID = result.Id.VideoId
	return
}
