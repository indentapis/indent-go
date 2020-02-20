package auditv1

import (
	"fmt"
	"net/url"
	"strings"

	auditpb "go.indent.com/indent-go/api/indent/audit/v1"
)

func TargetFromDSN(dsn string) (*auditpb.Target, error) {
	if u, err := url.Parse(dsn); err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %v", err)
	} else if tName := targetName(u); len(tName) != 32 {
		return nil, fmt.Errorf("target '%s' not valid", tName)
	} else {
		return &auditpb.Target{
			Name: tName,
		}, nil
	}
}

func targetName(u *url.URL) string {
	return strings.TrimPrefix(u.Path, "/v1/")
}
