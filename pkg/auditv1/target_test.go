package auditv1

import (
	"testing"

	"github.com/golang/protobuf/proto"

	auditpb "go.indent.com/indent-go/api/indent/audit/v1"
)

func TestTargetFromDSN(t *testing.T) {
	tests := []struct {
		dsn    string
		target *auditpb.Target
		fail   bool
	}{
		{
			dsn:    "https://audit.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			target: &auditpb.Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
		{
			dsn:    "https://test-user:test-password@audit.indentapis.com/v1/12951631f22a14ecd927d0002b91df93",
			target: &auditpb.Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
		{
			dsn:    "/v1/12951631f22a14ecd927d0002b91df93",
			target: &auditpb.Target{Name: "12951631f22a14ecd927d0002b91df93"},
		},
	}

	for i, test := range tests {
		if target, err := TargetFromDSN(test.dsn); err == nil && test.fail {
			t.Errorf("test %d failed: error should have occurred", i)
		} else if err != nil && !test.fail {
			t.Errorf("test %d failed: unexpected error: %v", i, err)
		} else if target == nil {
			t.Errorf("test %d failed: target can't be nil", i)
		} else if !proto.Equal(target, test.target) {
			t.Errorf("test %d failed: actual doesnt match expected", i)
		}
	}
}
