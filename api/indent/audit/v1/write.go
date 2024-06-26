package auditv1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
)

// Write event asynchronously to Audit API. Errors are logged including event.
func (c *Client) Write(event *Event) {
	if c.Debug {
		c.writeDebug(event)
		return
	}

	ts := ptypes.TimestampNow()
	go func() {
		// deep-copy event
		event, _ = proto.Clone(event).(*Event)

		// set timestamp to current time
		event.Timestamp = ts

		ctx := context.Background()
		if err := c.WriteEvents(ctx, []*Event{event}); err != nil {
			c.Log.Error("Failed writing event", zap.Error(err))
		} else {
			c.Log.Debug("successfully wrote event")
		}
	}()
}

// WriteDebug event asynchronously adds events to a list of DebugEvents.
func (c *Client) writeDebug(event *Event) {
	ts := ptypes.TimestampNow()
	// deep-copy event
	event, _ = proto.Clone(event).(*Event)

	// set timestamp to current time
	event.Timestamp = ts

	if c.DebugEvents == nil {
		c.DebugEvents = []*Event{}
	}

	c.DebugEvents = append(c.DebugEvents, event)
	c.Log.Debug("successfully wrote event", zap.String("event", event.Event))
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
