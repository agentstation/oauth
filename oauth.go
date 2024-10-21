package oauth

import (
	"github.com/agentstation/egothic"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	// GetState gets the state from the oauth session.
	GetState = egothic.GetState
	// SetState sets the state in the oauth session.
	SetState = egothic.SetState
	// GetProviderName gets the provider name from the oauth session.
	GetProviderName = egothic.GetProviderName
)

// SetStore sets the store for the oauth session.
func SetStore(store sessions.Store) {
	egothic.SetStore(store)
}

// Store returns the store for the oauth session.
func Store() sessions.Store {
	return egothic.Store()
}

// Begin starts the authentication process for a given provider.
func Begin(e echo.Context, opts ...Options) error {
	eGothicOpts := optionsToEgothicOptions(opts...)
	return egothic.BeginAuthHandler(e, eGothicOpts...)
}

// Complete completes the authentication process and returns the user.
func Complete(e echo.Context, opts ...Options) (User, error) {
	eGothicOpts := optionsToEgothicOptions(opts...)
	guser, err := egothic.CompleteUserAuth(e, eGothicOpts...)
	if err != nil {
		return User{}, err
	}
	user := User{
		RawData:           guser.RawData,
		Provider:          guser.Provider,
		Email:             guser.Email,
		Name:              guser.Name,
		FirstName:         guser.FirstName,
		LastName:          guser.LastName,
		NickName:          guser.NickName,
		Description:       guser.Description,
		UserID:            guser.UserID,
		AvatarURL:         guser.AvatarURL,
		Location:          guser.Location,
		AccessToken:       guser.AccessToken,
		AccessTokenSecret: guser.AccessTokenSecret,
		RefreshToken:      guser.RefreshToken,
		ExpiresAt:         guser.ExpiresAt,
		IDToken:           guser.IDToken,
	}
	return user, nil
}

// ProviderURL returns the authentication URL for the specified provider.
func ProviderURL(e echo.Context, opts ...Options) (string, error) {
	eGothicOpts := optionsToEgothicOptions(opts...)
	return egothic.GetAuthURL(e, eGothicOpts...)
}

// SetSession stores a key/value pair in the session.
func SetSession(e echo.Context, key string, value string) error {
	return egothic.StoreInSession(e, key, value)
}

// GetSession retrieves a value from the session by key.
// It returns an error if the key doesn't exist.
func GetSession(e echo.Context, key string) (string, error) {
	return egothic.GetFromSession(e, key)
}

// Logout invalidates a user session.
func Logout(e echo.Context) error {
	return egothic.Logout(e)
}
