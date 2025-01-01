package twitch

type RefreshingAuthProvider struct {
	store RefreshingAuthProviderStorer
}

func (p *RefreshingAuthProvider) GetAccessTokenByUserId(userId string) (string, error) {
	return "", nil
}

func NewRefreshingAuthProvider(store RefreshingAuthProviderStorer) (*RefreshingAuthProvider, error) {
	authProvider := &RefreshingAuthProvider{
		store,
	}
	return authProvider, nil
}
