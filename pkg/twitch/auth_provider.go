package twitch

type AuthProvider interface {
	GetAccessToken() (string, error)
}
