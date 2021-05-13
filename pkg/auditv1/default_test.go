package auditv1

import (
	"context"
	"testing"
)

func TestDefaultWrite(t *testing.T) {
	Init(testTarget)
	for _, e := range testEvents(50) {
		Write(e)
	}
}

func TestDefaultWriteEvents(t *testing.T) {
	Init(testTarget)
	events := testEvents(50)

	ctx := context.Background()
	if err := WriteEvents(ctx, events); err != nil {
		t.Error(err)
	}
}
