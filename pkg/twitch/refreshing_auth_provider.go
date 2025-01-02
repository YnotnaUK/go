package twitch

import (
	"fmt"
	"sync"
)

type RefreshingAuthProvider struct {
	accessTokens map[string]*AccessToken
	clientId     string
	clientSecret string
	mutex        *sync.RWMutex
	redirectURI  string
}

type RefreshingAuthProviderConfig struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
}

func (p *RefreshingAuthProvider) AddAccessToken(accessToken *AccessToken) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.accessTokens[accessToken.UserId] = accessToken
	return nil
}

func (p *RefreshingAuthProvider) AddAccessTokenFromFile(fileLocation string) error {
	accessToken, err := CreateAccessTokenFromJSONFile(fileLocation)
	if err != nil {
		return err
	}
	err = p.AddAccessToken(accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (p *RefreshingAuthProvider) GetAccessTokenByUserId(userId string) (*AccessToken, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	// Get access token
	accessToken, ok := p.accessTokens[userId]
	if !ok {
		return nil, fmt.Errorf("no access token for user id: %s", userId)
	}
	// Validate access token
	_, err := ValidateAccessToken(accessToken.AccessToken)
	if err != nil {
		return nil, err
	}
	// Return access token
	return accessToken, nil
}

func NewRefreshingAuthProvider(config RefreshingAuthProviderConfig) (*RefreshingAuthProvider, error) {
	// Check config
	if config.ClientId == "" {
		return nil, ErrBlankClientId
	}
	if config.ClientSecret == "" {
		return nil, ErrBlankClientSecret
	}
	if config.RedirectURI == "" {
		return nil, ErrBlankRedirectURI
	}
	// Create provider
	authProvider := &RefreshingAuthProvider{
		accessTokens: make(map[string]*AccessToken),
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		mutex:        &sync.RWMutex{},
		redirectURI:  config.RedirectURI,
	}
	return authProvider, nil
}
