package auditv1

import (
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestTargetFromDSN(t *testing.T) {
	tests := []struct {
		dsn    string
		target *Target
		fail   bool
	}{
		{
			dsn:    "https://write.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			target: &Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
		{
			dsn:    "https://test-user:test-password@write.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			target: &Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
		{
			dsn:    "/v1/12951631f22a14ecd927d0002b91df93",
			target: &Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			target, err := targetFromDSN(test.dsn)
			switch test.fail {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}
			assert.NotNil(t, target)
			assert.True(t, proto.Equal(target, test.target))
		})
	}
}

func TestEndpointFromDSN(t *testing.T) {
	tests := []struct {
		dsn      string
		endpoint string
		fail     bool
	}{
		{
			dsn:      "https://write.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			endpoint: "write.indentapis.com:443",
		},
		{
			dsn:      "https://test-user:test-password@write.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			endpoint: "write.indentapis.com:443",
		},
		{
			dsn:      "/v1/12951631f22a14ecd927d0002b91df93",
			endpoint: "write.indentapis.com:443",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			endpoint, err := endpointFromDSN(test.dsn)
			switch test.fail {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}
			assert.Equal(t, test.endpoint, endpoint)
		})
	}
}
