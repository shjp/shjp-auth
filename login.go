package auth

import (
	"encoding/json"

	"github.com/pkg/errors"

	core "github.com/shjp/shjp-core"
)

// UserLogin represents a login session
type UserLogin struct {
	user          core.User
	strategy      LoginStrategy
	sessionClient SessionClient
}

// NewLogin instantiates a new UserLogin
func NewLogin(u core.User, strategy LoginStrategy, sessionClient SessionClient) *UserLogin {
	return &UserLogin{
		user:          u,
		strategy:      strategy,
		sessionClient: sessionClient,
	}
}

// Login logs the user in
func (l *UserLogin) Login() (*UserSession, error) {
	u, err := l.strategy.Verify(l.user)
	if err != nil {
		return nil, errors.Wrap(err, "Login strategy failed")
	}

	if err = u.PopulatePermissions(); err != nil {
		return nil, errors.Wrap(err, "Error populating user permissions")
	}

	userBytes, err := json.Marshal(u)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling user object")
	}

	sessionKey := generateSessionKey()
	if err = l.sessionClient.Set(sessionKey, userBytes); err != nil {
		return nil, errors.Wrap(err, "Error setting the session key/value")
	}

	return &UserSession{Key: sessionKey, User: *u}, nil
}
