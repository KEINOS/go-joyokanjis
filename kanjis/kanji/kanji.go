package kanji

// ----------------------------------------------------------------------------
//  Constants
// ----------------------------------------------------------------------------

// Range of CJK Unified Ideographs including extension upto extension D.
const (
	// The first rune of CJK Unified Ideograph.
	minCJK = 0x4e00
	// The last rune of CJK Unified Ideograph (including extension A to D).
	maxCJK = 0x2b81d
)

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// IsCJK returns true if the given rune is in a range of CJK Unified Ideographs
// (including extensions upto Extension D).
func IsCJK(r rune) bool {
	return r >= minCJK && r <= maxCJK
}

// ----------------------------------------------------------------------------
//  Type: Kanji
// ----------------------------------------------------------------------------

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
