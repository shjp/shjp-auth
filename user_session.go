package auth

import core "github.com/shjp/shjp-core"

// UserSession represents the user's current session
type UserSession struct {
	Key       string `json:"key"`
	core.User `json:"user"`
}
