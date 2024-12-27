package twitch

import twitch_auth "github.com/ynotnauk/go/pkg/twitch/auth"

type Bot struct {
	auth twitch_auth.AuthProvider
}

func (b *Bot) Start() {

}

func NewBot() (*Bot, error) {
	bot := &Bot{}
	return bot, nil
}

// Creates a fully ready to use bot
func NewSimpleBot(clientId string) (*Bot, error) {
	authProvider, err := twitch_auth.NewStaticAuthProvider(clientId)
	if err != nil {
		return nil, err
	}
	bot := &Bot{
		auth: authProvider,
	}
	return bot, nil
}
