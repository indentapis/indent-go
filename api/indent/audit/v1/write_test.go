package auditv1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteEvents(t *testing.T) {
	client, events := testClient(t), testEvents(10)

	ctx := context.Background()
	assert.NoError(t, client.WriteEvents(ctx, events))
}

func TestWriteEventsWithDSN(t *testing.T) {
	client, events := testClientFromDSN(t), testEvents(10)

	ctx := context.Background()
	assert.NoError(t, client.WriteEvents(ctx, events))
}
