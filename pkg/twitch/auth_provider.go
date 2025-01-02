package twitch

type AuthProvider interface {
	GetAccessTokenByUserId(userId string) (*AccessToken, error)
	AddAccessToken(accessToken *AccessToken) error
	AddAccessTokenFromFile(fileLocation string) error
}
