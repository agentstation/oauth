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
	// NOTE: The email address is the most important field and so it is required.
	// If the email is not found return an error, we can't accept the user.
	if u.Email == "" {
		return errors.New("email not found")
	}

	// if the user first name is not found and the full name is not empty, set the first name and last name
	if u.FirstName == "" && u.Name != "" {
		// parse the user full name and set the user's first name as the parsed first name, even if the parsed first name is blank
		first, last := parseFullName(u.Name)
		u.FirstName = first

		// if the user last name is empty and the first name is not the same as the parsed last name, set the last name
		if u.LastName == "" && last != "" && u.FirstName != last {
			u.LastName = last
		}
	}

	// if the user first name is still empty, but the user nickname is shorter than the user name, use the user nickname as the user first name
	if u.FirstName == "" && u.NickName != "" && u.Name != "" && (len(u.NickName) < len(u.Name)) {
		u.FirstName = u.NickName
	}

	// if the user first name is still empty, and the user full name is not empty, use the user full name as the user first name
	if u.FirstName == "" && u.Name != "" {
		u.FirstName = u.Name
	}

	// if the user first name is still empty, but the user nickname is available, use the user nickname as the user first name
	if u.FirstName == "" && u.NickName != "" {
		u.FirstName = u.NickName
	}

	// if the user first name is still empty, use the user email as the user first name as a fallback of last resort
	if u.FirstName == "" {
		u.FirstName = u.Email
	}

	return nil
}

// parseFullName parses a full name and returns the first and last name.
func parseFullName(fullName string) (string, string) {
	// parse the full name and return the first and last name
	parsedName := names.ParseFullName(fullName)
	first := parsedName.First
	last := parsedName.Last

	// if the first name and last name are the same, return an empty first and last name
	if first == last {
		return "", ""
	}

	// if the first name is empty, return an empty first and last name
	if first == "" {
		return "", ""
	}

	// return the first and last name
	return first, last
}
