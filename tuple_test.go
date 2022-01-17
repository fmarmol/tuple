package tuple

import (
	"testing"

	"gotest.tools/assert"
)

func TestTuple2(t *testing.T) {
	tu := NewTuple2(1, "one")
	assert.Equal(t, 1, tu.First)
	assert.Equal(t, "one", tu.Second)
}
