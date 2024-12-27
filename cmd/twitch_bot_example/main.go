package main

import (
	"log"
	"os"

	"github.com/ynotnauk/go/pkg/dotenv"
	"github.com/ynotnauk/go/pkg/twitch"
)

func main() {
	// Load env
	dotenv.Load()
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
