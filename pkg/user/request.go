package user

// GithubLoginRequest represents request data for POST /login
// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps#2-users-are-redirected-back-to-your-site-by-github
// swagger:model
type GithubLoginRequest struct {
	// The Github oAuth code received as oAuth response
	// example: ABCDEF
	Code string `json:"code" binding:"required"`

	// The unguessable random string provided in the initial oAuth request
	// example: 12345
	State string `json:"state" binding:"required"`
}
