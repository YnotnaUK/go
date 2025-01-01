package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ynotnauk/go/pkg/twitch"
)

var scopes = []string{
	"analytics:read:extensions",
	"analytics:read:games",
	"bits:read",
	"channel:bot",
	"channel:manage:ads",
	"channel:read:ads",
	"channel:manage:broadcast",
	"channel:read:charity",
	"channel:edit:commercial",
	"channel:read:editors",
	"channel:manage:extensions",
	"channel:read:goals",
	"channel:read:guest_star",
	"channel:manage:guest_star",
	"channel:read:hype_train",
	"channel:manage:moderators",
	"channel:read:polls",
	"channel:manage:polls",
	"channel:read:predictions",
	"channel:manage:predictions",
	"channel:manage:raids",
	"channel:read:redemptions",
	"channel:manage:redemptions",
	"channel:manage:schedule",
	"channel:read:stream_key",
	"channel:read:subscriptions",
	"channel:manage:videos",
	"channel:read:vips",
	"channel:manage:vips",
	"clips:edit",
	"moderation:read",
	"moderator:manage:announcements",
	"moderator:manage:automod",
	"moderator:read:automod_settings",
	"moderator:manage:automod_settings",
	"moderator:read:banned_users",
	"moderator:manage:banned_users",
	"moderator:read:blocked_terms",
	"moderator:read:chat_messages",
	"moderator:manage:blocked_terms",
	"moderator:manage:chat_messages",
	"moderator:read:chat_settings",
	"moderator:manage:chat_settings",
	"moderator:read:chatters",
	"moderator:read:followers",
	"moderator:read:guest_star",
	"moderator:manage:guest_star",
	"moderator:read:moderators",
	"moderator:read:shield_mode",
	"moderator:manage:shield_mode",
	"moderator:read:shoutouts",
	"moderator:manage:shoutouts",
	"moderator:read:suspicious_users",
	"moderator:read:unban_requests",
	"moderator:manage:unban_requests",
	"moderator:read:vips",
	"moderator:read:warnings",
	"moderator:manage:warnings",
	"user:bot",
	"user:edit",
	"user:edit:broadcast",
	"user:read:blocked_users",
	"user:manage:blocked_users",
	"user:read:broadcast",
	"user:read:chat",
	"user:manage:chat_color",
	"user:read:email",
	"user:read:emotes",
	"user:read:follows",
	"user:read:moderated_channels",
	"user:read:subscriptions",
	"user:read:whispers",
	"user:manage:whispers",
	"user:write:chat",
}

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create Token Generator Config
	tokenGeneratorConfig := &twitch.TokenGeneratorConfig{
		ClientId:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		Port:         8080,
		RedirectURI:  os.Getenv("TWITCH_REDIRECT_URI"),
		Scopes:       scopes,
	}
	// Create generator
	generator, err := twitch.NewTokenGenerator(tokenGeneratorConfig)
	if err != nil {
		log.Panic(err)
	}
	// Start Generator
	generator.Run()
}
