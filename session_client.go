package auth

import "github.com/shjp/shjp-core/model"

// SessionClient provides the top level interface client for user session
type SessionClient interface {
	Get(key string) (*model.User, error)
	Set(model.User) (string, error)
}

// SessionClientOptions provides the interface to populate a SessionClient
type SessionClientOptions interface {
	NewClient() (SessionClient, error)
	String() string
}
