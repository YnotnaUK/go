package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ynotnauk/go/pkg/twitch"
)

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	twitchClientId := os.Getenv("TWITCH_CLIENT_ID")
	//twitchClientSecret := os.Getenv("TWITCH_CLIENT_SECRET")
	// Create Bot
	bot, err := twitch.NewSimpleBot(twitchClientId)
	if err != nil {
		log.Panic(err)
	}
	// Start Bot
	bot.Start()
}
