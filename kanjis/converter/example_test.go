package converter_test

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/KEINOS/go-joyokanjis/kanjis/converter"
)

func Example_shift_up_code_point() {
	tf := converter.New(
		// This function shifts the code point of the input rune by 1
		func(in rune) rune {
			return in + 1 // e.g. 'a' -> 'b'
		},
	)

	// Input reader
	input := "abcdefgァィゥェォ"
	sReader := strings.NewReader(input)

	// Output writer
	var bWriter bytes.Buffer

	// Convert
	tf.Convert(sReader, &bWriter)

	fmt.Println(bWriter.String())
	// Output: bcdefghアイウエオ
}

func Example_toUpper() {
	tf := converter.New(
		// This function converts the input rune to upper case
		func(in rune) rune {
			return unicode.ToUpper(in)
		},
	)

	// Input reader
	input := "Hello, World!"
	sReader := strings.NewReader(input)

	// Output writer
	var bWriter bytes.Buffer

	// Convert
	tf.Convert(sReader, &bWriter)

	fmt.Println(bWriter.String())
	// Output: HELLO, WORLD!
}
