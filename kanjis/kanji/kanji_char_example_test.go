package kanji_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Code-Hex/dd"
	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/MakeNowJust/heredoc"
)

// ============================================================================
//  KanjiChar type
// ============================================================================

// ----------------------------------------------------------------------------
//  KanjiChar.MarshalJSON()
// ----------------------------------------------------------------------------

func ExampleKanjiChar_MarshalJSON() {
	charRaw := '滞'
	charKanji := kanji.KanjiChar(charRaw)

	jsonChar, err := charKanji.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonChar))
	// Output:
	// "滞"
}

func ExampleKanjiChar_String() {
	charRaw := '滞'
	charKanji := kanji.KanjiChar(charRaw)

	fmt.Println("Raw:", charRaw)
	fmt.Println("Stringer:", charKanji) // It implements the fmt.Stringer interface
	fmt.Println("String:", charKanji.String())
	// Output:
	// Raw: 28382
	// Stringer: 滞
	// String: 滞
}

// ----------------------------------------------------------------------------
//  KanjiChar.UnmarshalJSON()
// ----------------------------------------------------------------------------

func ExampleKanjiChar() {
	// Note that each field is a KanjiChar type, whcih is a rune.
	type SampleStruct struct {
		Kanji1 kanji.KanjiChar `json:"kanji_1"`
		Kanji2 kanji.KanjiChar `json:"kanji_2"`
		Kanji3 kanji.KanjiChar `json:"kanji_3"`
	}

	// Note that "kanji_2" has 2 runes. KanjiChar will only take the first rune.
	jsonData := heredoc.Doc(`{
		"kanji_1":"忍",
		"kanji_2":"忍者",
		"kanji_3":"者"
	}`)

	var sampleObj SampleStruct

	// Parse JSON object and store the result to the pointer.
	err := json.Unmarshal([]byte(jsonData), &sampleObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dd.Dump(sampleObj))
	// Output:
	// kanji_test.SampleStruct{
	//   Kanji1: 24525,
	//   Kanji2: 24525,
	//   Kanji3: 64091,
	// }
}
