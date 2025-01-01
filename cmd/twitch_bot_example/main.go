package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ynotnauk/go/pkg/twitch"
)

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create store
	store, err := twitch.NewBotFilesystemStore("data/twitchbotstore")
	if err != nil {
		log.Fatal(err)
	}
	// Create bot
	bot, err := twitch.NewBot(store)
	if err != nil {
		log.Fatal(err)
	}
	// Start bot
	bot.Start()
}
