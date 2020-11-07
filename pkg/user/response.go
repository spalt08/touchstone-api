package user

import "jsbnch/pkg/model"

// LoginResponse represents response for POST /login
// swagger:model
type LoginResponse struct {
	// Logged user
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
