package twitch

type BotStorer interface {
	GetAccessTokenByUserId(userId string) (*AccessToken, error)
	UpdateAccessTokenForUserId(userId string, accessToken *AccessToken) (*AccessToken, error)
}
