/*
Package kanjis is a set of handy function to access the singleton kanji.Dict object of the embedded dictionary.

It provides functions to:

1. Convert old kanji (kyujitai, 旧字体, 旧漢字) to new kanji (shinjitai, 新字体, 新漢字).
1. Detect if the given character is a joyo kanji (常用漢字) from shinjitai (新字体・新漢字).

*/
//go:generate go run internal/converter.go
package kanjis

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"

	"github.com/KEINOS/go-joyokanjis/kanjis/internal/tool"
	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/pkg/errors"
	"github.com/tidwall/transform"
)

// Private global variables for singleton object.
var (
	// gzData is the embedded GZipped Gob encoded data.
	//
	//go:embed internal/gzgob/dict.gzip
	gzData []byte
	// kanjiDict is the dictionary of Joyo Kanjis.
	kanjiDict kanji.Dict
)

// ----------------------------------------------------------------------------
//  Initialization
// ----------------------------------------------------------------------------

func init() {
	// Extract and decode the embedded archived dictionary to kanjiDict when
	// the package is imported.
	if err := extractEmbeddedData(); err != nil {
		log.Fatal(err)
	}
}

// ----------------------------------------------------------------------------
//  Public functions
// ----------------------------------------------------------------------------

// FixRuneAsJoyo returns the Joyo Kanji if the given character is a registered
// Kyujitai (old kanji) and has a new kanji (shinjitai) in the dictionary.
//
// Otherwise, it returns the given character. Thus, ASCII and other non-Japanese
// characters are returned as is.
// Also note that kyujitai characters that do not have a shinjitai are returned
// as is as well.
func FixRuneAsJoyo(char rune) rune {
	return kanjiDict.FixAsJoyo(char)
}

// FixStringAsJoyo is similar to FixRuneAsJoyo but for string.
//
// If the input is larger than 320 Bytes, consider using FixFileAsJoyo() instead.
func FixStringAsJoyo(input string) string {
	if input == "" {
		return ""
	}

	inRune := []rune(input)
	for i, char := range inRune {
		inRune[i] = kanjiDict.FixAsJoyo(char)
	}

	return string(inRune)
}

// FixFileAsJoyo is similar to FixRuneAsJoyo but for large data such as files.
func FixFileAsJoyo(input io.Reader, output io.Writer) error {
	sc := bufio.NewScanner(spawnWorker(input))
	for sc.Scan() {
		if _, err := output.Write([]byte(sc.Text() + "\n")); err != nil {
			return errors.Wrap(err, "failed to write the output file")
		}
	}

	if err := sc.Err(); err != nil && !errors.Is(err, io.EOF) {
		return errors.Wrap(err, "failed to scan the input file")
	}

	return nil
}

// IsJoyokanji returns true if the given rune is a Joyo Kanji character.
func IsJoyokanji(char rune) bool {
	return kanjiDict.IsJoyokanji(char)
}

// LenDict returns the number of Joyo Kanjis registered in the dictionary.
func LenDict() int {
	return kanjiDict.LenJoyo()
}

// ----------------------------------------------------------------------------
//  Private functions
// ----------------------------------------------------------------------------

func extractEmbeddedData() error {
	// Read embedded gzipped data
	src := bytes.NewReader(gzData)
	// Extract and decode the embedded GZipped Gob encoded data and assign to kanjiDict
	err := tool.ExtractGzipGobToDict(src, &kanjiDict)

	return errors.Wrap(err, "failed to extract and decode the embedded GZipped Gob encoded data")
}

func spawnWorker(src io.Reader) io.Reader {
	br := bufio.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		char, _, err := br.ReadRune()
		if err != nil {
			return nil, errors.Wrap(err,
				"failed to read a line during transformation")
		}

		return []byte(string(FixRuneAsJoyo(char))), nil
	})
}
