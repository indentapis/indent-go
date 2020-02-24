package auditv1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
)

// Write event asynchronously to Audit API. Errors are logged including event.
func (c *Client) Write(event *Event) {
	ts := ptypes.TimestampNow()
	go func() {
		event.Timestamp = ts

		ctx := context.Background()
		if err := c.WriteEvents(ctx, []*Event{event}); err != nil {
			c.Log.Errorf("failed writing event: %v", err)
		} else {
			c.Log.Debug("successfully wrote event")
		}
	}()
}

// WriteEvents to Audit API. Failures contained in err will contain event.
func (c *Client) WriteEvents(ctx context.Context, events []*Event) (err error) {
	req := &WriteRequest{
		Target: c.Target,
		Async:  true,
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
