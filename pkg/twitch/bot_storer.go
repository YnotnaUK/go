package twitch

type BotStorer interface {
	GetAccessToken(userId string) (*AccessToken, error)
	CreateOrUpdateAccessToken(accessToken *AccessToken) (bool, error)
}
