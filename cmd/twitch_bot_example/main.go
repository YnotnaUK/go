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
	// Create auth provider
	_, err = twitch.NewRefreshingAuthProvider()
	if err != nil {
		log.Fatal(err)
	}
	// Create bot
	bot, err := twitch.NewBot()
	if err != nil {
		log.Fatal(err)
	}
	// Start bot
	bot.Start()
}
