package twitch

type Bot struct {
	auth AuthProvider
}

func (b *Bot) Start() {

}

func NewBot() (*Bot, error) {
	bot := &Bot{}
	return bot, nil
}

// Creates a fully ready to use bot
func NewSimpleBot(clientId string) (*Bot, error) {
	authProvider, err := NewRefreshingAuthProvider()
	if err != nil {
		return nil, err
	}
	bot := &Bot{
		auth: authProvider,
	}
	return bot, nil
}
