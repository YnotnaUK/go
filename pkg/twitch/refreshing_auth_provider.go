package twitch

type RefreshingAuthProvider struct {
	store RefreshingAuthProviderStore
}

func NewRefreshingAuthProvider(store RefreshingAuthProviderStore) (*RefreshingAuthProvider, error) {
	refreshingAuthProvider := &RefreshingAuthProvider{
		store: store,
	}
	return refreshingAuthProvider, nil
}
