package music

import (
	"bytes"
	"encoding/json"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type videoResponse struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

func GetVideoDownloadUrl(videoId string) (url string, err error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", videoId)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Println("ERROR: something wrong happened when running youtube-dl")
		return
	}

	var videoRes videoResponse
	println(&out)
	err = json.NewDecoder(&out).Decode(&videoRes)
	if err != nil {
		log.Println("ERROR: error occurred when decoding video response")
		return
	}
	err = json.NewDecoder(&out).Decode(&videoRes)
	if err != nil {
		log.Println("ERROR: error occurred when decoding video response")
		return
	}
	return videoRes.Formats[2].Url, nil
}
