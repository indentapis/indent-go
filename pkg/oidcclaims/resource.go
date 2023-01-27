package oidcclaims

import (
	"errors"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
)

var (
	// ErrNoStandardClaims is returned when no Standard Claims have been set.
	ErrNoStandardClaims = errors.New("no Standard Claims found")
)

// Resource returned based on a set of claims.
func Resource(claims *Claims) (*auditv1.Resource, error) {
	if claims == nil || claims.Standard == nil {
		return nil, ErrNoStandardClaims
	}

	return &auditv1.Resource{
		Id:          claims.Sub,
		DisplayName: displayName(claims),
		Email:       claims.Email,
	}, nil
}

func displayName(claims *Claims) (out string) {
	if claims.Name != "" {
		return claims.Name
	}

	nameParts := []string{claims.GivenName, claims.MiddleName, claims.FamilyName}
	for _, part := range nameParts {
		if part != "" {
			if out != "" {
				out += " "
			}
			out += part
		}
	}
	return
}
