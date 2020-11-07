package user

// LoginRequest represents request data for POST /login
// swagger:model
type LoginRequest struct {
	// Github access token
	// example: ABCDEF
	AccessToken string `json:"accessToken" binding:"required"`
}
