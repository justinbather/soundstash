package downloader

import (
	"os"

	"github.com/justinbather/soundstash/server/internal/core/ports"
)

type downloader struct{}

func New() ports.Downloader {
	return &downloader{}
}

func (*downloader) Download(url string) (*os.File, error) {
	return nil, nil
}

/*
	url := "https://soundcloud.com/bbbostonbun/missing-you?in=just-justin-tho/sets/chill-house"
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
*/
