package email

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	core "github.com/shjp/shjp-core"
)

// LoginStrategy implements auth.LoginStrategy for email login
type LoginStrategy struct {
	DaoURL string
}

// Verify verifies the login for the given user
func (s *LoginStrategy) Verify(u core.User) (*core.User, error) {
	emailUser := core.User{
		Email:    u.Email,
		Password: u.Password,
	}

	payload, err := json.Marshal(emailUser)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshaling payload")
	}
	log.Println("Searching for user:", string(payload))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Post(s.DaoURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "Error received from DAO server")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading the response body")
	}

	var users []core.User
	if err = json.Unmarshal(body, &users); err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling the response to users array")
	}

	if len(users) == 0 {
		return nil, errors.New("User not found")
	}

	if len(users) > 1 {
		return nil, errors.New("Multiple users found for the given credentials")
	}

	log.Println("[email.LoginStrategy] users =", users)

	return &users[0], nil
}
