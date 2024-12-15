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

	if u.FirstName != "" {
		return nil
	}

	u.setFirstAndLastName()
	return nil
}

// setFirstAndLastName attempts to set the first and last name using available information
// in the following priority order: full name parsing, nickname, full name, email
func (u *User) setFirstAndLastName() {
	if u.tryParseFullName() {
		return
	}

	u.useEmailAsFirstName()
}

func (u *User) tryParseFullName() bool {
	if u.Name == "" {
		return false
	}

	parsedName := names.ParseFullName(u.Name)

	if parsedName.First != "" {
		u.FirstName = parsedName.First
	} else {
		if !u.tryUseNickname() {
			u.FirstName = u.Name
		}
	}

	if u.LastName == "" && parsedName.Last != "" && parsedName.Last != u.FirstName {
		u.LastName = parsedName.Last
	}

	return u.FirstName != ""
}

func (u *User) tryUseNickname() bool {
	if u.NickName == "" {
		return false
	}

	// if the nickname is shorter than the name, use it as the first name
	if len(u.NickName) < len(u.Name) {
		u.FirstName = u.NickName
		return true
	}

	return false
}

func (u *User) useEmailAsFirstName() {
	u.FirstName = u.Email
}
