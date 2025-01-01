package twitch

type RefreshingAuthProvider struct{}

func (p *RefreshingAuthProvider) GetAccessToken() (string, error) {
	return "", nil
}

func NewRefreshingAuthProvider() (*RefreshingAuthProvider, error) {
	authProvider := &RefreshingAuthProvider{}
	return authProvider, nil
}
