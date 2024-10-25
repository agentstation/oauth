```sh
                          ___   __ _  _   _  _    _     
                         / _ \ / _` || | | || |_ | |__  
                        | (_) | (_| || |_| || __|| '_ \ 
                         \___/ \__,_| \__,_| \__||_| |_|
```

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# oauth

```go
import "github.com/agentstation/oauth"
```

## Index

- [Variables](<#variables>)
- [func Begin\(e echo.Context, opts ...Options\) error](<#Begin>)
- [func GetSession\(e echo.Context, key string\) \(string, error\)](<#GetSession>)
- [func Logout\(e echo.Context\) error](<#Logout>)
- [func ProviderURL\(e echo.Context, opts ...Options\) \(string, error\)](<#ProviderURL>)
- [func SetProviders\(providers ...goth.Provider\)](<#SetProviders>)
- [func SetSession\(e echo.Context, key string, value string\) error](<#SetSession>)
- [func SetStore\(store sessions.Store\)](<#SetStore>)
- [func Store\(\) sessions.Store](<#Store>)
- [type Options](<#Options>)
  - [func WithDebug\(\) Options](<#WithDebug>)
  - [func WithLogger\(logger \*log.Logger\) Options](<#WithLogger>)
- [type Provider](<#Provider>)
- [type ProviderConfigMap](<#ProviderConfigMap>)
  - [func NewProviderConfigMap\(\) ProviderConfigMap](<#NewProviderConfigMap>)
  - [func \(pc ProviderConfigMap\) Get\(provider Provider\) ProviderKeyValues](<#ProviderConfigMap.Get>)
  - [func \(pc ProviderConfigMap\) Set\(provider Provider, keyValues ...string\) error](<#ProviderConfigMap.Set>)
- [type ProviderKeyValues](<#ProviderKeyValues>)
- [type User](<#User>)
  - [func Complete\(e echo.Context, opts ...Options\) \(User, error\)](<#Complete>)
  - [func \(u \*User\) Validate\(\) error](<#User.Validate>)


## Variables

<a name="GetState"></a>

```go
var (
    // GetState gets the state from the oauth session.
    GetState = egothic.GetState
    // SetState sets the state in the oauth session.
    SetState = egothic.SetState
    // GetProviderName gets the provider name from the oauth session.
    GetProviderName = egothic.GetProviderName
)
```

<a name="Begin"></a>
## func [Begin](<https://github.com/agentstation/oauth/blob/master/oauth.go#L30>)

```go
func Begin(e echo.Context, opts ...Options) error
```

Begin starts the authentication process for a given provider.

<a name="GetSession"></a>
## func [GetSession](<https://github.com/agentstation/oauth/blob/master/oauth.go#L76>)

```go
func GetSession(e echo.Context, key string) (string, error)
```

GetSession retrieves a value from the session by key. It returns an error if the key doesn't exist.

<a name="Logout"></a>
## func [Logout](<https://github.com/agentstation/oauth/blob/master/oauth.go#L81>)

```go
func Logout(e echo.Context) error
```

Logout invalidates a user session.

<a name="ProviderURL"></a>
## func [ProviderURL](<https://github.com/agentstation/oauth/blob/master/oauth.go#L64>)

```go
func ProviderURL(e echo.Context, opts ...Options) (string, error)
```

ProviderURL returns the authentication URL for the specified provider.

<a name="SetProviders"></a>
## func [SetProviders](<https://github.com/agentstation/oauth/blob/master/providers.go#L118>)

```go
func SetProviders(providers ...goth.Provider)
```

SetProviders sets the goth oauth providers.

You can find the list of supported providers here: https://github.com/markbates/goth?tab=readme-ov-file#supported-providers

<a name="SetSession"></a>
## func [SetSession](<https://github.com/agentstation/oauth/blob/master/oauth.go#L70>)

```go
func SetSession(e echo.Context, key string, value string) error
```

SetSession stores a key/value pair in the session.

<a name="SetStore"></a>
## func [SetStore](<https://github.com/agentstation/oauth/blob/master/oauth.go#L20>)

```go
func SetStore(store sessions.Store)
```

SetStore sets the store for the oauth session.

<a name="Store"></a>
## func [Store](<https://github.com/agentstation/oauth/blob/master/oauth.go#L25>)

```go
func Store() sessions.Store
```

Store returns the store for the oauth session.

<a name="Options"></a>
## type [Options](<https://github.com/agentstation/oauth/blob/master/options.go#L10>)

Options is a function that configures the oauth package.

```go
type Options func(*oauthConfig)
```

<a name="WithDebug"></a>
### func [WithDebug](<https://github.com/agentstation/oauth/blob/master/options.go#L46>)

```go
func WithDebug() Options
```

WithDebug sets the debug mode for the oauth package.

<a name="WithLogger"></a>
### func [WithLogger](<https://github.com/agentstation/oauth/blob/master/options.go#L53>)

```go
func WithLogger(logger *log.Logger) Options
```

WithLogger sets the logger for the oauth package.

<a name="Provider"></a>
## type [Provider](<https://github.com/agentstation/oauth/blob/master/providers.go#L10>)

Provider represents an OAuth provider

```go
type Provider string
```

<a name="Amazon"></a>Provider constants

```go
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
```

<a name="ProviderConfigMap"></a>
## type [ProviderConfigMap](<https://github.com/agentstation/oauth/blob/master/providers.go#L80>)

ProviderConfigMap is a map of providers to configuration key values.

```go
type ProviderConfigMap map[Provider]ProviderKeyValues
```

<a name="NewProviderConfigMap"></a>
### func [NewProviderConfigMap](<https://github.com/agentstation/oauth/blob/master/providers.go#L83>)

```go
func NewProviderConfigMap() ProviderConfigMap
```

NewProviderConfigMap creates a new provider configuration map.

<a name="ProviderConfigMap.Get"></a>
### func \(ProviderConfigMap\) [Get](<https://github.com/agentstation/oauth/blob/master/providers.go#L98>)

```go
func (pc ProviderConfigMap) Get(provider Provider) ProviderKeyValues
```

Get returns the ProviderKeyValues for the given provider name.

<a name="ProviderConfigMap.Set"></a>
### func \(ProviderConfigMap\) [Set](<https://github.com/agentstation/oauth/blob/master/providers.go#L88>)

```go
func (pc ProviderConfigMap) Set(provider Provider, keyValues ...string) error
```

Set sets the ProviderKeyValues for the given provider name.

<a name="ProviderKeyValues"></a>
## type [ProviderKeyValues](<https://github.com/agentstation/oauth/blob/master/providers.go#L77>)

ProviderKeyValues is the configuration for a goth oauth provider.

```go
type ProviderKeyValues map[string]string
```

<a name="User"></a>
## type [User](<https://github.com/agentstation/oauth/blob/master/user.go#L11-L28>)

User represents an oauth user that has authenticated.

```go
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
```

<a name="Complete"></a>
### func [Complete](<https://github.com/agentstation/oauth/blob/master/oauth.go#L36>)

```go
func Complete(e echo.Context, opts ...Options) (User, error)
```

Complete completes the authentication process and returns the user.

<a name="User.Validate"></a>
### func \(\*User\) [Validate](<https://github.com/agentstation/oauth/blob/master/user.go#L31>)

```go
func (u *User) Validate() error
```

Validate validates the user

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->