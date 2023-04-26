package common

const (
	// EventApprove is the event for when a request is approved.
	EventApprove = "access/approve"

	// EventDeny is the event for when a request is denied.
	EventDeny = "access/deny"

	// EventGrant is the event produced when a permission is granted.
	EventGrant = "access/grant"

	// EventRevoke is the event produced when a permission is revoked.
	EventRevoke = "access/revoke"
)

const (
	// LabelExpires contains a Time specifying the expiration of a resource.
	LabelExpires = IndentPrefix + "time/expires"

	// LabelDuration contains a length of time.
	LabelDuration = IndentPrefix + "time/duration"

	// LabelIsIndefinite is a boolean label indicating whether a claim approves
	// indefinite access.
	LabelIsIndefinite = IndentPrefix + "time/indefinite"
)
