package twitch

import "errors"

var (
	ErrNoAuthProvider  error = errors.New("no auth provider provided")
	ErrNoStoreProvided error = errors.New("no store provided")
	ErrNoUserId        error = errors.New("no user id provided")
)

type Bot struct {
	authProvider AuthProvider
	store        BotStorer
	userId       string
}

type BotConfig struct {
	AuthProvider AuthProvider
	Store        BotStorer
	UserId       string
}

func (b *Bot) Start() error {
	return nil
}

func NewBot(config *BotConfig) (*Bot, error) {
	// Validate config
	if config.AuthProvider == nil {
		return nil, ErrNoAuthProvider
	}
	if config.Store == nil {
		return nil, ErrNoStoreProvided
	}
	if config.UserId == "" {
		return nil, ErrNoUserId
	}
	// Create bot
	bot := &Bot{
		authProvider: config.AuthProvider,
		store:        config.Store,
		userId:       config.UserId,
	}
	// Return bot
	return bot, nil
}
