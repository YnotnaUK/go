package twitch

type Bot struct{}

func (b *Bot) Start() error {
	return nil
}

func NewBot() (*Bot, error) {
	bot := &Bot{}
	return bot, nil
}
