// Package oidcclaims allows reading of claims including Standard Claims.
package oidcclaims

import (
	"encoding/json"
	"fmt"
	"reflect"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
)

// Claims about a User.
type Claims struct {
	*Standard
	All map[string]interface{}
}

// Resource created by using Claims.
func (c *Claims) Resource() (*auditv1.Resource, error) {
	return Resource(c)
}

// UnmarshalJSON decodes Claims.
func (c *Claims) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.All); err != nil {
		return fmt.Errorf("failed decoding claims into map: %w", err)
	} else if err = json.Unmarshal(data, &c.Standard); err != nil {
		return fmt.Errorf("failed decoding standard claims: %w", err)
	} else if isEmpty(*c.Standard) {
		c.Standard = nil
	}
	return nil
}

func isEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}
