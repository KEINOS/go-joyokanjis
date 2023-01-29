package kanji

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Dict is a map of Kanji objects. The key is the rune (int32) that represents
// the Kanji.
type Dict map[rune]Kanji

// FixAsJoyo converts the given Kanji to the Joyo Kanji if it is a KyuJitai.
// If the given Kanji is not a KyuJitai, then it returns as is.
func (d Dict) FixAsJoyo(kanji rune) rune {
	tmpKanji, ok := d[kanji]
	if !ok || !tmpKanji.IsKyuJitai {
		return kanji
	}

	//return []rune(tmpKanji.ShinJitai)[0]
	return rune(tmpKanji.ShinJitai)
}

// IsJoyokanji returns true if the given Kanji is a Joyo Kanji.
func (d Dict) IsJoyokanji(kanji rune) bool {
	if tmpKanji, ok := d[kanji]; ok {
		return rune(tmpKanji.ShinJitai) == kanji
	}

	return false
}

// IsKyuJitai returns true if the given Kanji is a KyuJitai (old kanji).
func (d Dict) IsKyuJitai(kanji rune) bool {
	if tmpKanji, ok := d[kanji]; ok {
		return rune(tmpKanji.KyuJitai) == kanji
	}

	return false
}

// KunYomi returns the KunYomi reading in hiragana of the given Kanji.
func (d Dict) KunYomi(kanji rune) []string {
	if tmpKanji, ok := d[kanji]; ok {
		return tmpKanji.Yomi.KunYomi
	}

	return []string{}
}

// LenJoyo returns the number of Joyo Kanjis registered in the dictionary.
// It does not count the KyuJitai kanji entry.
func (d Dict) LenJoyo() int {
	var count int

	for _, kanji := range d {
		if !kanji.IsKyuJitai {
			count++
		}
	}

	return count
}

// OnYomi returns the OnYomi reading in katakana of the given Kanji.
func (d Dict) OnYomi(kanji rune) []string {
	if tmpKanji, ok := d[kanji]; ok {
		return tmpKanji.Yomi.OnYomi
	}

	return []string{}
}

// Marshal returns the JSON encoding of the dictionary.
func (d Dict) Marshal() ([]byte, error) {
	byteJSON, err := json.Marshal(d.stripKyuJitai())

	return byteJSON, errors.Wrap(err, "failed to marshal the dictionary")
}

// MarshalIndent returns the indented JSON encoding of the dictionary.
func (d Dict) MarshalIndent(prefix string, indent string) ([]byte, error) {
	byteJSON, err := json.MarshalIndent(d.stripKyuJitai(), prefix, indent)

	return byteJSON, errors.Wrap(err, "failed to marshal the dictionary")
}

// stripKyuJitai removes the kujitai entry (element with kyujitai key) from the
// dictionary.
func (d Dict) stripKyuJitai() Dict {
	newDict := make(map[rune]Kanji, len(d))

	for keyKanji, valueKanji := range d {
		if valueKanji.IsKyuJitai {
			continue
		}

		newDict[keyKanji] = valueKanji
	}

	return newDict
}
