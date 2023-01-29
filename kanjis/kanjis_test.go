package kanjis

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJoyokanji(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		char         rune
		expectIsJoyo bool
	}{
		{'漢', true},
		{'漢', false},
		{'巣', true},
		{'巢', false},
		{'a', false},
		{'あ', false},
		{'ア', false},
	} {
		actualIsJoyo := IsJoyokanji(test.char)

		require.Equal(t, test.expectIsJoyo, actualIsJoyo)
	}
}

func TestFixStringAsJoyo_goroutine(t *testing.T) {
	t.Parallel()

	const input = "これは舊漢字です。"

	var result string
	var wg sync.WaitGroup

	// Run in goroutine
	wg.Add(1)

	go func() {
		defer wg.Done()

		result = FixStringAsJoyo(input)
	}()

	// Wait for goroutine to finish
	wg.Wait()

	require.Equal(t, "これは旧漢字です。", result)
}
