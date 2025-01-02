package twitch

type RefreshingAuthProvider struct {
	accessTokens map[string]AccessToken
	clientId     string
	clientSecret string
	redirectURI  string
}

type RefreshingAuthProviderConfig struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
}

func (p *RefreshingAuthProvider) GetAccessTokenByUserId(userId string) (string, error) {
	return "", nil
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
		accessTokens: make(map[string]AccessToken),
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		redirectURI:  config.RedirectURI,
	}
	return authProvider, nil
}
