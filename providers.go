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

// ProviderKeyValues is the configuration for a goth oauth provider.
type ProviderKeyValues map[string]string

// ProviderConfigMap is a map of providers to configuration key values.
type ProviderConfigMap map[string]ProviderKeyValues

// NewProviderConfigMap creates a new provider configuration map.
func NewProviderConfigMap() ProviderConfigMap {
	return make(ProviderConfigMap)
}

// Set sets the ProviderKeyValues for the given provider name.
func (pc ProviderConfigMap) Set(providerName string, keyValues ...string) {
	pc[providerName] = newProviderKeyValues(keyValues...)
}

// Get returns the ProviderKeyValues for the given provider name.
func (pc ProviderConfigMap) Get(providerName string) ProviderKeyValues {
	return pc[providerName]
}

// newProviderKeyValues creates a new provider configuration key values.
func newProviderKeyValues(keyValues ...string) ProviderKeyValues {
	keyValueMap := make(ProviderKeyValues)
	for i := 0; i < len(keyValues); i += 2 {
		if i+1 < len(keyValues) {
			keyValueMap[keyValues[i]] = keyValues[i+1]
		}
	}
	return keyValueMap
}
