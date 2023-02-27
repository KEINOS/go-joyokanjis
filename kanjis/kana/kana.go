/*
Package kana provides a type for Kana characters and functions for katakana and hiragana conversion.
*/
package kana

// codeDiff is the difference between hiragana and katakana code points in
// Unicode. Since the sequence of hiragana and katakana is relative, they can be
// converted to each other by adjusting the difference.
const codeDiff = 'ァ' - 'ぁ' // 0x60 = 0d96

// IsHiragana returns true if the rune is a Hiragana (hira-kana) that is convertable
// to Katakana.
//
// To determine Katakana that has no Hiragana equivalent, use `unicode.In(r, unicode.Hiragana)`
// instead.
func IsHiragana(r rune) bool {
	return (r >= 'ぁ' && r <= 'ゖ')
}

// IsKatakana returns true if the rune is a Katakana that is convertable to
// Hiragana.
//
// To determine Hiragana that has no Katakana equivalent, use `unicode.In(r, unicode.Katakana)`
// instead.
func IsKatakana(r rune) bool {
	return r >= 'ァ' && r <= 'ヶ'
}

// ToKatakana converts the given Hiragana rune to Katakana.
// If the given rune is not Hiragana convertable, then it returns as is.
func ToKatakana(r rune) rune {
	if IsHiragana(r) {
		return r + codeDiff
	}

	return r
}

// ToHiragana converts the given Katakana rune to Hiragana.
// If the given rune is not Katakana convertable, then it returns as is.
func ToHiragana(r rune) rune {
	if IsKatakana(r) {
		return r - codeDiff
	}

	return r
}
