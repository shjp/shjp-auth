package jwt

import (
	auth "github.com/shjp/shjp-auth"
)

// Options contains the option parameters for JWT
type Options struct {
}

// NewClient instantiates a new JWT client from the options
func (o *Options) NewClient() (auth.SessionClient, error) {
	return &Client{}, nil
}

// String stringifies the options object
func (o *Options) String() string {
	return ""
}
