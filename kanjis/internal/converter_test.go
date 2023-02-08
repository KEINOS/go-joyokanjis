package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/KEINOS/go-joyokanjis/kanjis/internal/tool"
	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

func Test_exitOnError_force_fail(t *testing.T) {
	require.PanicsWithError(t, "forced error", func() {
		exitOnError(errors.New("forced error"))
	}, "it should panic with the error reason")
}

func Test_main_golden(t *testing.T) {
	// JSON response from the test server
	responseJSON := heredoc.Doc(`{
		"134047": {
			"joyo_kanji": "𠮟",
			"yomi": {
				"on_yomi": [
					"シツ"
				],
				"kun_yomi": [
					"しか"
				],
				"example_yomi": [
					"しか-る"
				]
			},
			"raw_info": "𠮟\t\t5\t7S\t2010\tシツ、しか-る"
		}
	}`)

	// Backup before mocking
	backupAndDeferRestore(t)

	// Mock the global variables
	pathDirTmp := t.TempDir()
	urlDictSource = spawnTestServer(t, responseJSON)
	pathJSONInput = filepath.Join(pathDirTmp, "joyo2010.json")
	pathGobOutput = filepath.Join(pathDirTmp, "dict.gob")
	pathGzipOutput = filepath.Join(pathDirTmp, "dict.gzip")

	out := capturer.CaptureStdout(func() {
		require.NotPanics(t, func() {
			main()
		})
	})

	require.Contains(t, out, "Downloading JSON file...",
		"it should print the message downloading if the JSON file does not exist")

	require.FileExists(t, pathJSONInput)
	require.FileExists(t, pathGobOutput)
	require.FileExists(t, pathGzipOutput)

	// If the above test passed, the following tests must pass as well.
	require.True(t, fileExists(pathJSONInput))
	require.True(t, fileExists(pathGobOutput))
	require.True(t, fileExists(pathGzipOutput))

	// Check the content of the downloaded JSON file
	downloadedJSON, err := os.ReadFile(pathJSONInput)
	require.NoError(t, err)

	require.Equal(t, strings.TrimSpace(responseJSON), strings.TrimSpace(string(downloadedJSON)))

	// Check the archived Gob file
	var kanjiDict kanji.Dict

	ptrFieGzipOutput, err := os.Open(pathGzipOutput)
	require.NoError(t, err, "failed to open the Gzip file")

	defer ptrFieGzipOutput.Close()

	err = tool.ExtractGzipGobToDict(ptrFieGzipOutput, &kanjiDict)
	require.NoError(t, err, "failed to extract the Gzip file and assign to the Dict")

	// Check the parsed data
	require.True(t, kanjiDict.IsJoyoKanji('𠮟'))
}

func Test_downloadDictJSON(t *testing.T) {
	// Backup before mocking the global variables
	backupAndDeferRestore(t)

	pathDirTmp := t.TempDir()
	oldURLDictSource := urlDictSource

	t.Run("invalid URL (empty URL)", func(t *testing.T) {
		urlDictSource = "" // Mock the global variable

		err := downloadDictJSON(pathDirTmp)

		require.Error(t, err, "invalid/empty URL should return an error")
		require.Contains(t, err.Error(), "failed to download the JSON dictionary",
			"it should contain the error reason")
	})

	t.Run("invalid path (empty path)", func(t *testing.T) {
		// Ensure the global variable is not empty
		urlDictSource = oldURLDictSource

		err := downloadDictJSON("") // set output path to empty

		require.Error(t, err, "empty path should return an error")
		require.Contains(t, err.Error(), "failed to create a file to save the JSON dictionary",
			"it should contain the error reason")
	})
}

// ----------------------------------------------------------------------------
//  Helper functions
// ----------------------------------------------------------------------------

// Backup global variables and restore them after the test.
func backupAndDeferRestore(t *testing.T) {
	t.Helper()

	oldURLDictSource := urlDictSource
	oldPathJSONInput := pathJSONInput
	oldPathGobOutput := pathGobOutput
	oldPathGzipOutput := pathGzipOutput

	t.Cleanup(func() {
		urlDictSource = oldURLDictSource
		pathJSONInput = oldPathJSONInput
		pathGobOutput = oldPathGobOutput
		pathGzipOutput = oldPathGzipOutput
	})
}

// Starts a new Server and returns it URL.
func spawnTestServer(t *testing.T, responseJSON string) string {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, responseJSON)
	}))

	t.Cleanup(func() {
		ts.Close()
	})

	return ts.URL
}
