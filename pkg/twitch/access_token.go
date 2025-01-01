package twitch

type AccessToken struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int      `json:"expires_in"`
	Login        string   `json:"login_id"`
	RefreshToken string   `json:"refresh_token"`
	Scopes       []string `json:"scopes"`
	UserId       string   `json:"user_id"`
}
