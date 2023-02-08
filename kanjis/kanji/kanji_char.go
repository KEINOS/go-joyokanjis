package kanji

import "github.com/goccy/go-json"

// ----------------------------------------------------------------------------
//  Type: KanjiChar
// ----------------------------------------------------------------------------

// KanjiChar is a rune that represents a Kanji.
// It is used to make rune readable in JSON format. It implements the Stringer,
// Marshaler and Unmarshaler interface.
type KanjiChar rune

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// String returns the string representation of the KanjiChar.
func (k KanjiChar) String() string {
	return string(k)
}

// MarshalJSON returns the JSON representation of the KanjiChar.
// It is a Marshaler interface implementation.
func (k KanjiChar) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(k))
}

// UnmarshalJSON unmarshals the JSON representation of the KanjiChar.
func (k *KanjiChar) UnmarshalJSON(b []byte) error {
	// Remove the double quotes
	b = b[1 : len(b)-1]

	// Convert to rune
	*k = KanjiChar([]rune(string(b))[0])

	return nil
}
