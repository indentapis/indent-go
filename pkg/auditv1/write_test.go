package auditv1

import (
	"context"
	"testing"
)

func TestWriteEvents(t *testing.T) {
	client, events := testClient(t), testEvents(100)

	ctx := context.Background()
	if err := client.WriteEvents(ctx, events); err != nil {
		t.Error(err)
	}
}
