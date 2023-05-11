package cliutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	"go.indent.com/indent-go/pkg/testutils"
)

func TestSelect(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// generate 50 resources
	resources := make([]*auditv1.Resource, 50)
	for i := 0; i < len(resources); i++ {
		resources[i] = testutils.Resource()
	}

	s, err := NewSelect(resources)
	r.NotEmpty(s)
	r.NoError(err)

	// check resources
	for i := 0; i < len(resources); i++ {
		r.Equal(resources[i], s.items[i])
	}
}

func TestSelectEmpty(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	s, err := NewSelect([]*auditv1.Event{})
	r.Empty(s)
	r.Error(err)
}
