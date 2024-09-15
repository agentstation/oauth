package oauth

import (
	"github.com/markbates/goth"
)

// SetProviders sets the goth oauth providers.
//
// You can find the list of supported providers here: https://github.com/markbates/goth?tab=readme-ov-file#supported-providers
func SetProviders(providers ...goth.Provider) {
	goth.UseProviders(providers...)
}
