package twitch

type Bot struct {
	authProvider AuthProvider
}

type BotConfig struct {
	AuthProvider AuthProvider
}

func (b *Bot) Start() error {
	return nil
}

func NewBot(config *BotConfig) (*Bot, error) {
	bot := &Bot{
		authProvider: config.AuthProvider,
	}
	return bot, nil
}
