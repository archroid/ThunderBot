package searchservice

import (
	"archroid/ElProfessorBot/static"
	"fmt"

	"github.com/sarulabs/di/v2"
	"google.golang.org/api/youtube/v3"
)

func GetVideoID(query string, container di.Container) (youtubeID string, err error) {
	searchService := container.Get(static.DiYoutubeSearch).(*youtube.SearchService)

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
