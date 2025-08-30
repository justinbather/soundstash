package main

import (
	"os"
	"os/exec"
)

func main() {
	log := newLogger()
	defer log.Sync()

	url := "https://soundcloud.com/bbbostonbun/missing-you?in=just-justin-tho/sets/chill-house"
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
}
