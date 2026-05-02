package auth

// OAuth2Token is the bearer token that protects every Garmin Connect API
// request. Tokens carry a TTL (ExpiresIn seconds) plus a longer-lived
// RefreshTokenExpiresIn so callers can persist them between runs.
type OAuth2Token struct {
	Scope                 string `json:"scope"`
	Jti                   string `json:"jti"`
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

// oauth1Token is the intermediate token Garmin returns from
// `oauth-service/oauth/preauthorized`. It is not exposed to callers because
// it is consumed immediately to mint an OAuth2 token.
type oauth1Token struct {
	Token       string
	TokenSecret string
}
