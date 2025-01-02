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
	// Create auth provider config
	authConfig := twitch.RefreshingAuthProviderConfig{
		ClientId:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("TWITCH_REDIRECT_URI"),
	}
	// Create auth provider
	_, err = twitch.NewRefreshingAuthProvider(authConfig)
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
