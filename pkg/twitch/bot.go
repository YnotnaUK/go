package twitch

type TwitchBot struct{}

func NewSimpleBot() (*TwitchBot, error) {
	bot := &TwitchBot{}
	return bot, nil
}
