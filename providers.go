package oauth

import "github.com/markbates/goth"

// Provider represents an OAuth provider
type Provider string

// Provider constants
const (
	Amazon          Provider = "amazon"
	Apple           Provider = "apple"
	Auth0           Provider = "auth0"
	AzureAD         Provider = "azuread"
	BattleNet       Provider = "battlenet"
	Bitbucket       Provider = "bitbucket"
	Box             Provider = "box"
	Dailymotion     Provider = "dailymotion"
	Deezer          Provider = "deezer"
	DigitalOcean    Provider = "digitalocean"
	Discord         Provider = "discord"
	Dropbox         Provider = "dropbox"
	EveOnline       Provider = "eveonline"
	Facebook        Provider = "facebook"
	Fitbit          Provider = "fitbit"
	Gitea           Provider = "gitea"
	GitHub          Provider = "github"
	GitLab          Provider = "gitlab"
	Google          Provider = "google"
	GPlus           Provider = "gplus"
	Heroku          Provider = "heroku"
	Instagram       Provider = "instagram"
	Intercom        Provider = "intercom"
	Kakao           Provider = "kakao"
	LastFM          Provider = "lastfm"
	Line            Provider = "line"
	LinkedIn        Provider = "linkedin"
	Mastodon        Provider = "mastodon"
	Meetup          Provider = "meetup"
	MicrosoftOnline Provider = "microsoftonline"
	Naver           Provider = "naver"
	Nextcloud       Provider = "nextcloud"
	Okta            Provider = "okta"
	OneDrive        Provider = "onedrive"
	OpenIDConnect   Provider = "openidconnect"
	Patreon         Provider = "patreon"
	PayPal          Provider = "paypal"
	Salesforce      Provider = "salesforce"
	SeaTalk         Provider = "seatalk"
	Shopify         Provider = "shopify"
	Slack           Provider = "slack"
	SoundCloud      Provider = "soundcloud"
	Spotify         Provider = "spotify"
	Steam           Provider = "steam"
	Strava          Provider = "strava"
	Stripe          Provider = "stripe"
	TikTok          Provider = "tiktok"
	Twitch          Provider = "twitch"
	Twitter         Provider = "twitter"
	TwitterV2       Provider = "twitterv2"
	Typetalk        Provider = "typetalk"
	Uber            Provider = "uber"
	VK              Provider = "vk"
	WeCom           Provider = "wecom"
	WePay           Provider = "wepay"
	Xero            Provider = "xero"
	Yahoo           Provider = "yahoo"
	Yammer          Provider = "yammer"
	Yandex          Provider = "yandex"
	Zoom            Provider = "zoom"
)

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

// SetProviders sets the goth oauth providers.
//
// You can find the list of supported providers here: https://github.com/markbates/goth?tab=readme-ov-file#supported-providers
func SetProviders(providers ...goth.Provider) {
	goth.UseProviders(providers...)
}
