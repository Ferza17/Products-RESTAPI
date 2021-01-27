package customers

type CustomerLoginResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	Exp         int64  `json:"expiration"`
}

type CustomerRefreshTokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
