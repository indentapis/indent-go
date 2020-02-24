package auditv1

import (
	"fmt"
	"net/url"
	"strings"
)

func TargetFromDSN(dsn string) (*Target, error) {
	if u, err := url.Parse(dsn); err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %v", err)
	} else if tName := targetName(u); len(tName) != 32 {
		return nil, fmt.Errorf("target '%s' not valid", tName)
	} else {
		return &Target{
			Name: tName,
		}, nil
	}
}

func targetName(u *url.URL) string {
	return strings.TrimPrefix(u.Path, "/v1/")
}
