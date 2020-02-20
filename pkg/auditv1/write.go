package auditv1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"

	auditpb "go.indent.com/indent-go/api/indent/audit/v1"
)

// Write event asynchronously to Audit API. Errors are logged including event.
func (c *Client) Write(event *auditpb.Event) {
	ts := ptypes.TimestampNow()
	go func() {
		event.Timestamp = ts

		ctx := context.Background()
		if err := c.WriteEvents(ctx, true, []*auditpb.Event{event}); err != nil {
			c.Log.Errorf("failed writing event: %v", err)
		} else {
			c.Log.Debug("successfully wrote event")
		}
	}()
}

// WriteEvents to Audit API. Failures contained in err will contain event.
func (c *Client) WriteEvents(ctx context.Context, async bool, events []*auditpb.Event) (err error) {
	req := &auditpb.WriteRequest{
		Target: c.Target,
		Async:  async,
		Events: events,
	}

	if _, err = c.Audit.Write(ctx, req); err != nil {
		err = &Error{
			Message: fmt.Sprintf("failed to write audit event: %v", err),
			Events:  events,
		}
	}
	return
}
