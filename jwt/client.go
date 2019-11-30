package jwt

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/shjp/shjp-core/model"
)

// Client is the JWT client agent
type Client struct {
}

// Get returns the payload from the JWT
func (c *Client) Get(key string) (*model.User, error) {
	token, err := jwtGo.ParseWithClaims(key, &userClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, errors.New("Error parsing with claims")
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, errors.New("Error casting userClaims from token")
	}

	if !token.Valid {
		return nil, errors.New("Token is not valid")
	}

	return &claims.User, nil
}

// Set for JWT simply returns the token (stateless)
func (c *Client) Set(user model.User) (string, error) {
	claims := userClaims{
		User: user,
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", errors.Wrap(err, "Error signing JWT string from token")
	}
	return ss, nil
}
