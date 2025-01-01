package twitch

type Bot struct {
	store BotStorer
}

func (b *Bot) Start() {}

func NewBot(store BotStorer) (*Bot, error) {
	bot := &Bot{
		store,
	}
	return bot, nil
}
