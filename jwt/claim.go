package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shjp/shjp-core/model"
)

var signingKey = []byte("foobar")

type userClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}
