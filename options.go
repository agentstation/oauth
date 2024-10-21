package oauth

import (
	"log"

	"github.com/agentstation/egothic"
)

// Options is a function that configures the oauth package.
type Options func(*oauthConfig)

type oauthConfig struct {
	debug  bool
	logger *log.Logger
}

func (c *oauthConfig) apply(opts ...Options) {
	for _, opt := range opts {
		opt(c)
	}
}

func newConfig(opts ...Options) *oauthConfig {
	c := &oauthConfig{}
	c.apply(opts...)
	return c
}

func (c *oauthConfig) toEgothicOptions() []egothic.Options {
	eGothicOpts := []egothic.Options{}
	if c.debug {
		eGothicOpts = append(eGothicOpts, egothic.WithDebug())
	}
	if c.logger != nil {
		eGothicOpts = append(eGothicOpts, egothic.WithLogger(c.logger))
	}
	return eGothicOpts
}

func optionsToEgothicOptions(opts ...Options) []egothic.Options {
	config := newConfig(opts...)
	return config.toEgothicOptions()
}

// WithDebug sets the debug mode for the oauth package.
func WithDebug() Options {
	return func(c *oauthConfig) {
		c.debug = true
	}
}

// WithLogger sets the logger for the oauth package.
func WithLogger(logger *log.Logger) Options {
	return func(c *oauthConfig) {
		c.logger = logger
	}
}
