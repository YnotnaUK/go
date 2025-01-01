package twitch

type Bot struct{}

func (b *Bot) Start() {}

func NewBot() (*Bot, error) {
	bot := &Bot{}
	return bot, nil
}
