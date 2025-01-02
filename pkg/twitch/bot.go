package twitch

import "fmt"

type Bot struct {
	authProvider AuthProvider
	userId       string
}

type BotConfig struct {
	AuthProvider AuthProvider
	UserId       string
}

func (b *Bot) Start() error {
	return nil
}

func NewBot(config *BotConfig) (*Bot, error) {
	bot := &Bot{
		authProvider: config.AuthProvider,
		userId:       config.UserId,
	}
	bot.authProvider.AddAccessTokenFromFile("data/access_token.json")
	fmt.Println(config.AuthProvider.GetAccessTokenByUserId(bot.userId))
	return bot, nil
}
