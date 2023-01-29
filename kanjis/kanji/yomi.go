package kanji

// Yomi is a struct that represents the Yomi (reading) of a Joyo Kanji.
type Yomi struct {
	// OnYomi is the list of OnYomi readings of the Kanji.
	OnYomi []string `json:"on_yomi,omitempty"`
	// KunYomi is the list of KunYomi readings of the Kanji.
	KunYomi []string `json:"kun_yomi,omitempty"`
	// ExampleYomi is the list of example readings of the Kanji.
	ExampleYomi []string `json:"example_yomi,omitempty"`
}
