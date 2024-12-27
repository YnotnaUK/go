package twitch_auth

type StaticAuthProvider struct {
	clientId string
}

func (a *StaticAuthProvider) GetAccessToken() (string, error) {
	return "fakeaccesstoken", nil
}

func NewStaticAuthProvider(clientId string) (AuthProvider, error) {
	authProvider := &StaticAuthProvider{
		clientId: clientId,
	}
	return authProvider, nil
}
