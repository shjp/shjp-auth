package auth

import core "github.com/shjp/shjp-core"

// LoginStrategy provides the strategy for a login method
type LoginStrategy interface {
	Verify(core.User) (*core.User, error)
}
