package twitch

type RefreshingAuthProviderFilesystemStore struct{}

func NewRefreshingAuthProviderFilesystemStore() (*RefreshingAuthProviderFilesystemStore, error) {
	store := &RefreshingAuthProviderFilesystemStore{}
	return store, nil
}
