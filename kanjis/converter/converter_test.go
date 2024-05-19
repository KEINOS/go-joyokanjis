package converter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConverter_Convert_nil_input(t *testing.T) {
	tf := New(
		func(in rune) rune {
			return in // do nothing
		},
	)

	t.Run("nil_input_should_error", func(t *testing.T) {
		var bWriter bytes.Buffer

		err := tf.Convert(nil, &bWriter)

		require.Error(t, err)
		require.Contains(t, err.Error(), "input or output is nil",
			"error message should contain the error reason")
	})

	t.Run("nil_output_should_error", func(t *testing.T) {
		sReader := strings.NewReader("test")
		err := tf.Convert(sReader, nil)

		require.Error(t, err)
		require.Contains(t, err.Error(), "input or output is nil",
			"error message should contain the error reason")
	})
}
