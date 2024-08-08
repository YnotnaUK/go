package twitch

type RefreshingAuthProviderStore interface {
	GetTokensForUserId(userId string) (*UserTokens, error)
}
