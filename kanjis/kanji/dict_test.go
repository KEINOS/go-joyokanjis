package kanji

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDict_input_is_nil(t *testing.T) {
	dictTest, err := NewDict(nil)

	require.Error(t, err, "nil input should return an error")
	assert.Contains(t, err.Error(), "failed to unmarshal JSON to Dict",
		"it should contain the error reason")
	require.Nil(t, dictTest, "the returned dictionary should be nil on error")
}
