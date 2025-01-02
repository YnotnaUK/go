package twitch

type RefreshingAuthProvider struct {
}

func (p *RefreshingAuthProvider) GetAccessTokenByUserId(userId string) (string, error) {
	return "", nil
}

func NewRefreshingAuthProvider() (*RefreshingAuthProvider, error) {
	authProvider := &RefreshingAuthProvider{}
	return authProvider, nil
}
