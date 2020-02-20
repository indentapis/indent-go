package auditv1

import (
	"context"
	"testing"
)

func init() {
	Init(testTarget)
}

func TestDefaultWrite(t *testing.T) {
	for _, e := range testEvents(50) {
		Write(e)
	}
}

func TestDefaultWriteEvents(t *testing.T) {
	events := testEvents(50)

	ctx := context.Background()
	if err := WriteEvents(ctx, true, events); err != nil {
		t.Error(err)
	}
}
