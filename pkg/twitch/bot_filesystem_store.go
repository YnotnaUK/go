package twitch

type BotFilesystemStore struct {
	storePath string
}

func (s *BotFilesystemStore) GetAccessTokenByUserId(userId string) (*AccessToken, error) {
	return nil, nil
}

func (s *BotFilesystemStore) UpdateAccessTokenForUserId(userId string, accessToken *AccessToken) (*AccessToken, error) {
	return nil, nil
}

func NewBotFilesystemStore(storeLocation string) (*BotFilesystemStore, error) {
	store := &BotFilesystemStore{
		storePath: storeLocation,
	}
	return store, nil
}
