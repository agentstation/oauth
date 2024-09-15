package oauth

import (
	"errors"
	"time"

	names "github.com/vinser/parse-full-name"
)

// User represents an oauth user that has authenticated.
type User struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

// Validate validates the user
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email not found")
	}

	if u.Name != "" && (u.FirstName == "" || u.LastName == "") {
		parsedName := names.ParseFullName(u.Name)
		u.FirstName = parsedName.First
		u.LastName = parsedName.Last
	}

	if u.FirstName == "" {
		return errors.New("first name not found")
	}

	return nil
}
