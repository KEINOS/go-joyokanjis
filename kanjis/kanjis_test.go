package kanjis

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// ----------------------------------------------------------------------------
//	init()
// ----------------------------------------------------------------------------

func Test_init_fail(t *testing.T) {
	oldGzData := gzData

	defer func() {
		gzData = oldGzData
	}()

	gzData = nil

	require.PanicsWithError(t,
		"initilization failed in package kanjis: failed to extract and decode the embedded GZipped Gob encoded data: failed to create a gzip reader: EOF",
		func() {
			initialize()
		},
	)
}

// ----------------------------------------------------------------------------
//	FixFileAsJoyo()
// ----------------------------------------------------------------------------

func TestFixFileAsJoyo_out_file_is_dir(t *testing.T) {
	t.Parallel()

	input := strings.NewReader("これは舊漢字です。")

	output, err := os.Open(t.TempDir())
	require.NoError(t, err, "failed to open temp dir")

	defer output.Close()

	err = FixFileAsJoyo(input, output)

	require.Error(t, err,
		"nil output file should return an error")
	assert.Contains(t, err.Error(), "failed to write the output file",
		"it should contain the error reason")
}

func TestFixFileAsJoyo_out_file_is_nil(t *testing.T) {
	t.Parallel()

	input := strings.NewReader("これは舊漢字です。")

	err := FixFileAsJoyo(input, nil)

	require.Error(t, err,
		"nil output file should return an error")
	assert.Contains(t, err.Error(), "input or output is nil",
		"it should contain the error reason")
}

func TestFixFileAsJoyo_fail_during_read(t *testing.T) {
	t.Parallel()

	// Create a dummy reader that returns an error on the 2nd read
	input := &DummyReader{
		Count:        0,
		ErrorOnCount: 1,
	}

	var output bytes.Buffer

	err := FixFileAsJoyo(input, &output)

	require.Error(t, err,
		"it should fail during the scan")
	assert.Contains(t, err.Error(), "failed to scan the input file",
		"it should contain the error reason")
}

// ----------------------------------------------------------------------------
//  FixRuneAsJoyo()
// ----------------------------------------------------------------------------

// This test detects if the given kanji is not in the Joyo Kanji list but has a
// new kanji form. Meaning, not in Joyo Kanji but is in kanji.NonJoyoOld2NewMap.
//
// To add a new kanji, edit the file `kanji/non_joyo_old2new_map.go`.
func TestFixRuneAsJoyo_new_old_comparison(t *testing.T) {
	pathFileData := filepath.Join("testdata", "shin_kyu.txt")

	ptrFile, err := os.Open(pathFileData)
	require.NoError(t, err, "failed to open test data file")

	defer ptrFile.Close()

	scanner := bufio.NewScanner(ptrFile)

	oldKanjis := []string{} // List of old kanjis
	newKanjis := []string{} // List of new kanjis

	count := 0

	for scanner.Scan() {
		count++

		if count%2 == 0 {
			lines := strings.Split(scanner.Text(), ",")
			oldKanjis = append(oldKanjis, lines[1:]...)

			continue
		}

		lines := strings.Split(scanner.Text(), ",")
		newKanjis = append(newKanjis, lines[1:]...)
	}

	require.Equal(t, len(oldKanjis), len(newKanjis),
		"oldKanjis and newKanjis should have the same length")

	for index, oldKanji := range oldKanjis {
		inputRune := []rune(oldKanji)[0]
		expect := []rune(newKanjis[index])[0]

		actual := FixRuneAsJoyo(inputRune) // Test fix

		assert.Equal(t, expect, actual,
			"old kanji %s (%d) should be converted to new kanji %s",
			string(inputRune), inputRune,
			string(expect))
	}
}

// ----------------------------------------------------------------------------
//  FixStringAsJoyo()
// ----------------------------------------------------------------------------

func TestFixStringAsJoyo_goroutine(t *testing.T) {
	t.Parallel()

	const (
		input         = "これは舊漢字です。"
		expect        = "これは旧漢字です。"
		numGoroutines = 3
	)

	var wg sync.WaitGroup

	// Spawn goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		// Run in goroutine
		go func(t *testing.T) {
			defer wg.Done()

			// Test
			assert.Equal(t, expect, FixStringAsJoyo(input),
				"FixStringAsJoyo() did not return the expected result")
		}(t)
	}

	// Wait for goroutine to finish
	wg.Wait()
}

// ----------------------------------------------------------------------------
//  IsJoyoKanji()
// ----------------------------------------------------------------------------

func TestIsJoyoKanji(t *testing.T) {
	t.Parallel()

	for index, test := range []struct {
		char         rune
		expectIsJoyo bool
	}{
		{'漢', true},  // new kanji
		{'漢', false}, // old kanji
		{'巣', true},  // new kanji
		{'巢', false}, // old kanji
		{'a', false}, // ASCII
		{'あ', false}, // Hiragana
		{'ア', false}, // Katakana
	} {
		expect := test.expectIsJoyo
		actual := IsJoyoKanji(test.char)

		require.Equal(t, expect, actual,
			"test #%d failed: IsJoyoKanji(%q) = %t; want %t", index, test.char, actual, expect)
	}
}

// ----------------------------------------------------------------------------
//  Miscellanous
// ----------------------------------------------------------------------------

// This test is to check if the kanjiDict singleton has the correct range of keys.
//
// If this test fails, it means that the imported joyo-kanji diciotnary has been
// updated. In that case, the constants such as kanji.minCJK and kanji.maxCJK
// needs to be modified accordingly.
func Test_range_keys_dict(t *testing.T) {
	expectLowestKey := rune(0x4e00)
	expectHighestKey := rune(0x20b9f)

	keys := maps.Keys(kanjiDict)
	slices.Sort(keys)

	actualLowestKey := keys[0]
	t.Logf("Lowest key: %x", actualLowestKey)

	actualHighestKey := keys[len(keys)-1]
	t.Logf("Highest key: %x", actualHighestKey)

	require.Equal(t, expectLowestKey, actualLowestKey, "Lowest key is not correct")
	require.Equal(t, expectHighestKey, actualHighestKey, "Highest key is not correct")

	// Check if the keys are in range of kanji.IsCJK function
	require.True(t, kanji.IsCJK(actualLowestKey), "Lowest key is not in range of kanji.IsCJK")
	require.True(t, kanji.IsCJK(actualHighestKey), "Highest key is not in range of kanji.IsCJK")
}

// ============================================================================
//  Helper functions/types
// ============================================================================

// DummyReader is a dummy reader that returns an error on the specified read
// count. Use this in cases where you want a delayed error.
type DummyReader struct {
	Count        int
	ErrorOnCount int
}

// Read implements the io.Reader interface.
func (r *DummyReader) Read(p []byte) (n int, err error) {
	r.Count++
	if r.Count == r.ErrorOnCount {
		return 0, os.ErrInvalid
	}

	return len(p), nil
}
