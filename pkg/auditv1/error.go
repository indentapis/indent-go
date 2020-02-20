package auditv1

import (
	"encoding/json"
	"fmt"

	auditpb "go.indent.com/indent-go/api/indent/audit/v1"
)

// Error encountered when writing an Event. Contains original Event for posterity.
type Error struct {
	// Message relating to failure.
	Message string

	// Events that failed to be written.
	Events []*auditpb.Event
}

// Error returns the reason for the failure as well as the original Event.
func (e *Error) Error() string {
	eventData, _ := json.Marshal(e.Events)
	return fmt.Sprintf("%s, event:\n[%s]", e.Message, eventData)
}
