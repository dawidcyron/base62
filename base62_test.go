package base62

import (
	"testing"

	"gotest.tools/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "2lz", ToBase62(10663))
}

func TestDecode(t *testing.T) {
	assert.Equal(t, 36, ToBase10("a"))
}
