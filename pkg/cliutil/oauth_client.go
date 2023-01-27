

package cliutil

import (
	"golang.org/x/oauth2"
)

var (
	oauthProduction = &oauth2.Config{
		Endpoint: oauthEndpoint,
		Scopes:   oauthScopes,
	}

	oauthStaging = &oauth2.Config{
		Endpoint: oauthEndpoint,
		Scopes:   oauthScopes,
	}
)
