package auth

import "github.com/shjp/shjp-core/model"

// UserSession represents the user's current session
type UserSession struct {
	Key        string `json:"key"`
	model.User `json:"user"`
}
