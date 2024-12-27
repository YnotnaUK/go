package twitch_auth

type AuthProvider interface {
	GetAccessToken() (string, error)
	//GetLoginAndAccessToken() (string, string, error)
}
