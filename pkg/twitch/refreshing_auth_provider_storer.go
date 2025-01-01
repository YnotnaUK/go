package twitch

type RefreshingAuthProviderStorer interface {
	GetAccessTokenByUserId(userId string) (*AccessToken, error)
	UpdateAccessTokenForUserId(userId string, accessToken *AccessToken) (*AccessToken, error)
}
