package oauthutil

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestLogin(t *testing.T) {
	if clientID, ok := os.LookupEnv("TEST_OKTA_CLIENT_ID"); ok {
		if secretID, ok := os.LookupEnv("TEST_OKTA_SECRET_ID"); ok {
			options := NewLoginOptions()
			options.OAuth = &oauth2.Config{
				ClientID:     clientID,
				ClientSecret: secretID,
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://sso.indent.com/oauth2/default/v1/authorize",
					TokenURL: "https://sso.indent.com/oauth2/default/v1/token",
				},
				Scopes: []string{"openid"},
			}

			code, err := Login(options)
			assert.NoError(t, err)

			_, err = options.OAuth.Exchange(context.Background(), code)
			assert.NoError(t, err)
		} else {
			t.Skip("TEST_OKTA_SECRET_ID must be provided")
		}
	} else {
		t.Skip("TEST_OKTA_CLIENT_ID must be provided")
	}
}
