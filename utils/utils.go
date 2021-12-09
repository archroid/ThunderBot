package utils

import (
	"errors"
	"time"
)

func VideoDurationValid(videoDuration time.Duration) (err error) {
	if videoDuration.Minutes() > 10 {
		err = errors.New("video is more than 10 minutes long")
	}
	return
}
