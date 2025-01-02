package twitch

type AuthProvider interface {
	GetAccessTokenByUserId(userId string) (string, error)
	AddAccessToken(accessToken *AccessToken) error
	AddAccessTokenFromFile(accessToken *AccessToken) error
}
