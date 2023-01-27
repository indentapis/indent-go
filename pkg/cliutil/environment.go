package cliutil

import (
	"golang.org/x/oauth2"
)

var (
	oauthEndpoint = oauth2.Endpoint{
		AuthURL:  "https://sso.indent.com/oauth2/default/v1/authorize",
		TokenURL: "https://sso.indent.com/oauth2/default/v1/token",
	}
	oauthScopes = []string{"openid", "email", "profile"}
)

var (
	// envProduction is the environment for production.
	envProduction = &Environment{
		Name:   "default",
		Target: "platform.indentapis.com:443",
		OAuth:  oauthProduction,
	}

	// envStaging is the environment for staging.
	envStaging = &Environment{
		Name:   "staging",
		Target: "platform.marble.indent.services:443",
		OAuth:  oauthStaging,
	}
)

// Environment specifies the Indent environment being connected to.
type Environment struct {
	// Name of the environment.
	Name string

	// Target is the gRPC endpoint dialed to connect.
	Target string

	// OAuth contains the configuration used to authenticate with OAuth2.
	OAuth *oauth2.Config
}
