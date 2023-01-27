package oauthutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"

	"go.indent.com/indent-go/pkg/oidcclaims"
)

const (
	pathUserInfo = "/v1/userinfo"
)

// UserInfo requests current information about the user.
func UserInfo(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*oidcclaims.Claims, error) {
	if config == nil {
		return nil, ErrMissingOAuthConfig
	} else if token == nil {
		return nil, ErrMissingOAuthToken
	}

	endpoint, err := endpointFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("couldn't determine baseURL: %w", err)
	}

	userinfoURL := endpoint + pathUserInfo
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userinfoURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed constructing userinfo request %q: %w", userinfoURL, err)
	}

	client := oauth2.NewClient(ctx, config.TokenSource(ctx, token))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed calling userinfo endpoint %q: %w", userinfoURL, err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading body of response from endpoint %q: %w", userinfoURL, err)
	} else if err = resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed closing while reading body of response from endpoint %q: %w", userinfoURL, err)
	}

	claims := new(oidcclaims.Claims)
	if err = json.Unmarshal(data, claims); err != nil {
		return nil, fmt.Errorf("failed decoding claims from endpoint %q: %w", userinfoURL, err)
	}
	return claims, err
}
