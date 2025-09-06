
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <Instagram Reel URL>")
		return
	}

	reelURL := ""
	cmd := exec.Command("yt-dlp", "-o", "reel.mp4", reelURL)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error downloading reel:", err)
		return
	}

	fmt.Println("Reel downloaded successfully → reel.mp4")
}
