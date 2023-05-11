package testutils

import (
	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	"go.indent.com/indent-go/pkg/rand"
)

// Resource returns a random Resource.
//
//nolint:gomnd // used for testing
func Resource() *auditv1.Resource {
	altIDs := make([]string, rand.Intn(10))
	for i := 0; i < len(altIDs); i++ {
		altIDs[i] = rand.Str(8)
	}

	numLabels := rand.Intn(10)
	labels := make(map[string]string, numLabels)
	for i := 0; i < numLabels; i++ {
		labels[rand.Str(8)] = rand.Str(15)
	}
	labels["example"] = rand.Str(15)

	// use both email and ID as identifiers
	id, email := rand.Str(8), rand.Str(22)
	if randIdentifier := rand.Intn(100); randIdentifier <= 30 {
		email = ""
	} else if randIdentifier >= 75 {
		id = ""
	}

	return &auditv1.Resource{
		Id:          id,
		Email:       email,
		DisplayName: rand.Str(22),
		AltIds:      altIDs,
		Kind:        rand.Str(8),
		Labels:      labels,
	}
}
