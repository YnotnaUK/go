package twitch

type AuthProvider interface {
	GetAccessTokenByUserId(userId string) (string, error)
}
