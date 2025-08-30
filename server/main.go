package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	url := "https://soundcloud.com/just-justin-tho/sets/chill-house"
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
