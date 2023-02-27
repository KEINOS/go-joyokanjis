/*
Package kanjis is a set of handy function to access the singleton kanji.Dict object of the embedded dictionary.

It provides functions to:

1. Convert old kanji (kyujitai, 旧字体, 旧漢字) to new kanji (shinjitai, 新字体, 新漢字).
1. Detect if the given character is a joyo kanji (常用漢字) from shinjitai (新字体・新漢字).
1. Search for the readings (読み, yomi) of the given kanji.

*/
//go:generate go run internal/converter.go
package kanjis

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/KEINOS/go-joyokanjis/kanjis/converter"
	"github.com/KEINOS/go-joyokanjis/kanjis/internal/tool"
	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/pkg/errors"
)

// Private global variables for singleton object.
var (
	// gzData is the embedded GZipped Gob encoded data.
	//
	//go:embed internal/gzgob/dict.gzip
	gzData []byte
	// kanjiDict is the singleton object that holds the Joyo Kanji dictionary.
	kanjiDict kanji.Dict
	// ignoreList
	ignoreList map[rune]interface{}
)

// ----------------------------------------------------------------------------
//  Initialization
// ----------------------------------------------------------------------------

func init() {
	// Spawn the singleton object of Joyo Kanji dictionary when the package is
	// imported.
	initialize()
}

func initialize() {
	// Extract and decode the embedded archived dictionary to kanjiDict.
	if err := extractEmbeddedData(); err != nil {
		panic(errors.Wrap(err, "initilization failed in package kanjis"))
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
	if ignoreList == nil {
		return kanjiDict.FixAsJoyo(char)
	}

	if _, ok := ignoreList[char]; ok {
		return char
	}

	return kanjiDict.FixAsJoyo(char)
}

// FixStringAsJoyo is similar to FixRuneAsJoyo but for string.
//
// If the input is larger than 320 Bytes, consider using FixFileAsJoyo() instead.
func FixStringAsJoyo(input string) string {
	inRune := []rune(input)

	for i, char := range inRune {
		inRune[i] = FixRuneAsJoyo(char)
	}

	return string(inRune)
}

// FixFileAsJoyo is similar to FixRuneAsJoyo but for file.
func FixFileAsJoyo(input io.Reader, output io.Writer) error {
	if input == nil || output == nil {
		return errors.New("input or output is nil")
	}

	tf := converter.New(func(in rune) rune {
		return FixRuneAsJoyo(in)
	})

	err := tf.Convert(input, output)

	return errors.Wrap(err, "failed to convert the input to the output")
}

// Ignore adds the given characters to the ignore list. These characters will be
// ignored when converting old kanji (kyujitai) to new kanji (shinjitai).
func Ignore(char ...rune) {
	if ignoreList == nil {
		ignoreList = make(map[rune]interface{})
	}

	for _, c := range char {
		ignoreList[c] = nil
	}
}

// IsJoyoKanji returns true if the given rune is a Joyo Kanji character.
func IsJoyoKanji(char rune) bool {
	return kanjiDict.IsJoyoKanji(char)
}

// IsKyuJitai returns true if the given rune is a registered Kyujitai (old kanji)
// character which contains a new kanji (shinjitai) in the dictionary.
func IsKyuJitai(char rune) bool {
	return kanjiDict.IsKyuJitai(char)
}

// LenDict returns the number of Joyo Kanjis registered in the dictionary.
func LenDict() int {
	return kanjiDict.LenJoyo()
}

// ResetIgnore clears the ignore list.
func ResetIgnore() {
	ignoreList = nil
}

// ----------------------------------------------------------------------------
//  Private functions
// ----------------------------------------------------------------------------

// extractEmbeddedData extracts the embedded GZipped Gob encoded dictionary and
// sets the decoded data to kanjiDict object as a singleton.
func extractEmbeddedData() error {
	// Read embedded gzipped data
	src := bytes.NewReader(gzData)

	// Extract and decode the embedded GZipped Gob encoded data and assign to
	// kanjiDict
	return errors.Wrap(tool.ExtractGzipGobToDict(src, &kanjiDict),
		"failed to extract and decode the embedded GZipped Gob encoded data")
}
