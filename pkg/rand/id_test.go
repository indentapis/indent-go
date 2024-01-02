package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDNum(t *testing.T) {
	id := IDNum()
	assert.NotZero(t, id)

	id2 := IDNum()
	assert.NotZero(t, id2)
	assert.NotEqual(t, id, id2)
}
