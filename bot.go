
package main

import (
	"fmt"
	"os"
	"os/exec"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := : ""
	bot , err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Error creating bot:", err)
		return
	}
	bot.debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		
		fmt.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

	url := update.Message.Text
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Downloading your reel...")
	bot.Send(msg)
	filename := fmt.Sprintf("reel_%d.mp4", time.Now().Unix())
	cmd := exec.Command("yt-dlp", "-o", filename, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Failed to download reel.")
		bot.Send(msg)
		continue
	}
	video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(filename))
	bot.Send(video)
	os.Remove(filename)
	}
}
	
