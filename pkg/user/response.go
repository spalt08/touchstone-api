package user

import "touchstone-api/pkg/model"

// GithubLoginResponse represents response for POST /login
// swagger:model
type GithubLoginResponse struct {
	// Logged user
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
