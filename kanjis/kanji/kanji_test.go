package kanji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test detects whether the key of NonJoyoOld2NewMap is in the range of IsCJK.
func Test_NonJoyoOld2NewMap_in_range(t *testing.T) {
	for key, val := range NonJoyoOld2NewMap {
		assert.True(t, IsCJK(key),
			"NonJoyoOld2NewMap key %s (%q) is not in range of IsCJK", string(key), key)
		assert.True(t, IsCJK(val),
			"NonJoyoOld2NewMap value %s (%q) is not in range of IsCJK", string(val), val)
	}
}
