package auth

import (
	core "github.com/shjp/shjp-core"
	"github.com/shjp/shjp-core/model"
)

// LoginStrategy provides the strategy for a login method
type LoginStrategy interface {
	Verify(core.User) (*model.User, error)
}
