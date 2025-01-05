package twitch

type BotStorer interface {
	GetAccessTokenForUserId(userId string) (*AccessToken, error)
	CreateOrUpdateAccessToken(accessToken *AccessToken) (bool, error)
}
