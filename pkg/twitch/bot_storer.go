package twitch

type BotStorer interface {
	GetAccessToken(userId string) (*AccessToken, error)
	CreateAccessToken(accessToken *AccessToken) (bool, error)
}
