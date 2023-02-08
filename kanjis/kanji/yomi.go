package kanji

import "github.com/KEINOS/go-joyokanjis/kanjis/kana"

// ----------------------------------------------------------------------------
//  Type: Yomi
// ----------------------------------------------------------------------------

// Yomi is a struct that represents the Yomi (reading, "読み") of a Kanji.
type Yomi struct {
	// OnYomi is the list of "On" readings (on-yomi, "音読み") of the Kanji.
	// Which is the reading derived from the Chinese pronunciations
	OnYomi []kana.Kanas `json:"on_yomi,omitempty"`
	// KunYomi is the list of "Kun" readings (kun-yomi, "訓読み") of the Kanji.
	// Which is the original, indigenous Japanese readings.
	KunYomi []kana.Kanas `json:"kun_yomi,omitempty"`
	// ExampleYomi is the list of example readings of the Kanji.
	ExampleYomi []string `json:"example_yomi,omitempty"`
}
