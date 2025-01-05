package twitch

import (
	"fmt"
)

type BotMemoryStore struct {
	accessTokens map[string]*AccessToken
}

func (s *BotMemoryStore) GetAccessTokenForUserId(userId string) (*AccessToken, error) {
	accessToken, ok := s.accessTokens[userId]
	if !ok {
		return nil, fmt.Errorf("no access token for user id: %s", userId)
	}
	return accessToken, nil
}
func (s *BotMemoryStore) CreateOrUpdateAccessToken(accessToken *AccessToken) (bool, error) {
	s.accessTokens[accessToken.UserId] = accessToken
	return true, nil
}

func NewBotMemoryStore() (*BotMemoryStore, error) {
	store := &BotMemoryStore{
		accessTokens: make(map[string]*AccessToken),
	}
	return store, nil
}
