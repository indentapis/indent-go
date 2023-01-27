package oauthutil

import (
	"fmt"
	"net/url"
	"path"

	"golang.org/x/oauth2"
)

func endpointFromConfig(config *oauth2.Config) (string, error) {
	endpoint := config.Endpoint.AuthURL
	if endpoint == "" {
		endpoint = config.Endpoint.TokenURL
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("couldn't determine Authorization base: %w", err)
	}
	u.Path = path.Dir(path.Dir(u.Path))
	return u.String(), nil
}
