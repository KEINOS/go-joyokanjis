package kanji

import "github.com/Code-Hex/dd"

// Kanji is a struct that represents a Joyo Kanji.
type Kanji struct {
	// Yomi is the reading info of the Kanji.
	Yomi Yomi `json:"yomi,omitempty"`
	// ShinJitai is the Joyo Kanji in new kanji form.
	ShinJitai KanjiChar `json:"joyo_kanji,omitempty"`
	// KyuJitai is the old kanji form which is mapped to the shinjitai.
	KyuJitai KanjiChar `json:"kyu_jitai,omitempty"`
	// IsKyuJitai is true if the map key is a KyuJitai.
	IsKyuJitai bool `json:"-"`
}

// String returns the structure of the Kanji object as a readable string (not JSON).
func (k Kanji) String() string {
	return dd.Dump(k)
}
