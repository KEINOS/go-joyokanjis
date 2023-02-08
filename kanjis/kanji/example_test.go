package kanji_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
)

// This example describes how to parse the JSON data to the kanji.Dict type and
// how to use the dictionary.
//
// For the JSON format see: https://gist.github.com/KEINOS/fb660943484008b7f5297bb627e0e1b1#format
func Example() {
	// Sample JSON data.
	// In this exmple only two kanji are registered. Note that "滞"(28382)
	// contains old kanji "滯"(28399) in "kyu_jitai" element.
	sampleJSON := `{
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
		},
		"28382": {
			"joyo_kanji": "滞",
			"kyu_jitai": "滯",
			"yomi": {
				"on_yomi": ["タイ"],
				"kun_yomi": ["とどこお"],
				"example_yomi": ["とどこお-る"]
			},
			"raw_info": "滞\t滯\t13\t7S\t\tタイ、とどこお-る"
		}
	}`

	// Create a new dictionary by unmarshaling the JSON dictionary.
	// If an element contains "kyu_jitai"(old kanji), then a new element for that
	// old kanji is also added to the dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Check if "滞"(28382, new kanji) is a joyo kanji.
	fmt.Println("Is 滞 joyo?:", tmpDict.IsJoyoKanji('滞'))
	// Print the on-yomi and kun-yomi of "滞"(28382, new kanji).
	fmt.Println("Reading of 滞:", "On yomi:", tmpDict.OnYomi('滞'), "Kun yomi:", tmpDict.KunYomi('滞'))

	// Since "滞"(28382) contains an old kanji "滯"(28399), a new entry with the
	// old kanji is also mapped to the dictionary automatically.
	//
	// Let's check if "滯"(28399, old kanji) is a joyo kanji (expecting false).
	fmt.Println("Is 滯 joyo?:", tmpDict.IsJoyoKanji('滯'))

	// For more examples of Dict type methods, see the godoc of Dict type.

	// Output:
	// Is 滞 joyo?: true
	// Reading of 滞: On yomi: [タイ] Kun yomi: [とどこお]
	// Is 滯 joyo?: false
}

// ============================================================================
//  Public functions
// ============================================================================

// ----------------------------------------------------------------------------
//  IsCJK()
// ----------------------------------------------------------------------------

func ExampleIsCJK() {
	for _, test := range []struct {
		input  rune
		expect bool
	}{
		// CJK
		{input: '一', expect: true},
		{input: '漢', expect: true},
		{input: '巣', expect: true},
		{input: '巢', expect: true},
		// Non-CJK
		{input: 'あ', expect: false},
		{input: 'ア', expect: false},
		{input: 'a', expect: false},
		{input: '+', expect: false},
	} {
		expect := test.expect
		actual := kanji.IsCJK(test.input)

		if expect != actual {
			log.Fatalf("%s expected to be %v but got %v", string(test.input), expect, actual)
		}
	}

	fmt.Println("OK")
	// Output: OK
}
