package auditv1

import (
	"fmt"
	"net/url"
	"strings"
)

func TargetFromDSN(dsn string) (*Target, error) {
	expectedTargetNameLen := 32
	if u, err := url.Parse(dsn); err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %w", err)
	} else if tName := targetName(u); len(tName) != expectedTargetNameLen {
		return nil, fmt.Errorf("target '%s' not valid", tName)
	} else {
		return &Target{
			Name: tName,
		}, nil
	}
}

func EndpointFromDSN(dsn string) (string, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return "", fmt.Errorf("failed to parse dsn: %w", err)
	}

	port := "443"
	if u.Scheme == "http" {
		port = "80"
	} else if u.Port() != "" {
		port = u.Port()
	}
	host := u.Hostname()
	if host == "" {
		host = defaultAuditHost
	}
	return host + ":" + port, nil
}

func targetName(u *url.URL) string {
	return strings.TrimPrefix(u.Path, "/v1/")
}
