package cliutil

import (
	"golang.org/x/oauth2"
)

const (
	productionClientID = "0oart83xeb1AKLXTG357"
	stagingClientID    = "0oartbh1tcDfIRZS9357"
)

var (
	oauthProduction = &oauth2.Config{
		ClientID: productionClientID,
		Endpoint: oauthEndpoint,
		Scopes:   oauthScopes,
	}

	oauthStaging = &oauth2.Config{
		ClientID: stagingClientID,
		Endpoint: oauthEndpoint,
		Scopes:   oauthScopes,
	}
)
