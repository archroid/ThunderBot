package youtubemusic

import (
	"errors"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	utils "archroid/ElProfessorBot/utils"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	ytdl "github.com/kkdai/youtube/v2/downloader"

	ytmeta "github.com/kkdai/youtube/v2"
)

const AudioPath = "audio-cache"

var ytID string

func Download(youtubeID string) (title string, err error) {
	os.Mkdir(AudioPath, 0777)

	client := ytmeta.Client{
		HTTPClient: http.DefaultClient,
	}
	video, err := client.GetVideo(youtubeID)
	if err != nil {
		log.Error("Failed to get video info: " + err.Error())
		return "", errors.New("video ID is invalid")
	}
	title = video.Title
	if err := utils.VideoDurationValid(video.Duration); err != nil {
		return "", err
	}

	format, audioFormatFound := findAudioFormat(video.Formats)

	if audioFormatFound {
		err = handleAudioFormat(client, video, format)
	} else {
		err = handleVideoFormat(client, video, format)
	}

	if err != nil {
		return "", err
	}

	ytID = youtubeID
	return title, nil
}

func findAudioFormat(formats ytmeta.FormatList) (*ytmeta.Format, bool) {
	if format, found := filterFormatListByMime(formats, "audio/mp4"); found {

		utils.LogInfo("audio_format_found", utils.KVs("id", ytID, "event", "ytmp3"))
		return format, true

	} else if format, found := filterFormatListByMime(formats, "mp4"); found {
		utils.LogInfo("mp4_format_found", utils.KVs("id", ytID, "event", "ytmp3"))
		return format, false
	}

	defaultFormat := formats[0]
	utils.LogError("default_video_format", utils.KVs("url", defaultFormat.URL, "event", "ytmp3"))
	return &defaultFormat, false
}

func filterFormatListByMime(formats ytmeta.FormatList, query string) (*ytmeta.Format, bool) {
	for _, format := range formats {
		if strings.ContainsAny(format.MimeType, query) {
			return &format, true
		}
	}
	return nil, false
}

func handleVideoFormat(client ytmeta.Client, vid *ytmeta.Video, format *ytmeta.Format) error {
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Println("ffmpeg not found")
		return err
	}

	youtubeID := vid.ID
	videoFname := filepath.Join(AudioPath, youtubeID+".mp4")
	mp3Fname := PathToAudio(youtubeID)

	// Cleanup resulting video file downloaded by downloadByFormat
	defer os.Remove(videoFname)

	err = downloadByFormat(client, vid, format, videoFname)
	if err != nil {
		return err
	}
	utils.LogInfo("video_ready", utils.KVs("name", vid.Title, "youtubeID", youtubeID, "event", "ytmp3"))

	cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", videoFname, "-vn", mp3Fname)
	if err := cmd.Run(); err != nil {
		utils.LogError("failed audio_extraction", utils.KVs("err", err.Error(), "youtubeID", youtubeID, "event", "ytmp3"))
		return err
	}

	utils.LogInfo("audio_extracted", utils.KVs("filename", mp3Fname, "youtubeID", youtubeID, "event", "ytmp3"))

	return nil
}

func handleAudioFormat(client ytmeta.Client, vid *ytmeta.Video, format *ytmeta.Format) error {

	mp3Fname := PathToAudio(vid.ID + ".mp3")

	err := downloadByFormat(client, vid, format, mp3Fname)

	if err != nil {
		return err
	}

	utils.LogInfo("audio_downloaded", utils.KVs("filename", mp3Fname, "youtubeID", vid.ID, "event", "ytmp3"))
	return nil
}

func downloadByFormat(client ytmeta.Client, vid *ytmeta.Video, format *ytmeta.Format, fname string) error {

	ctx := context.Background()

	downloader := ytdl.Downloader{
		client,
		"",
	}
	err := downloader.Download(ctx, vid, format, fname)
	if err != nil {
		utils.LogError("video_download_failed", utils.KVs("name", vid.Title, "youtubeID", vid.ID, "event", "ytmp3", "err", err.Error()))
		return err
	}

	return nil
}

func PathToAudio(youtubeID string) string {
	return filepath.Join(AudioPath, youtubeID)
}
