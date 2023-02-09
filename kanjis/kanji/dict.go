package kanji

import (
	"encoding/json"

	"github.com/KEINOS/go-joyokanjis/kanjis/kana"
	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
//  Type: Dict
// ----------------------------------------------------------------------------

// Dict is a map of Kanji objects that works as a hash table.
// The key is the rune (int32) that represents the Kanji.
type Dict map[rune]Kanji

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// NewDict parses the JSON data of Joyo Kanji dictionary to kanji.Dict object.
//
// If the registered Joyo Kanji has an old kanji (kyu jitai), an alias key will
// be added to the dictionary to speed up the search.
//
// See the following URL for the format of the JSON byte array:
//
//	https://gist.github.com/KEINOS/fb660943484008b7f5297bb627e0e1b1#format
func NewDict(jsonDict []byte) (*Dict, error) {
	tmpDict := new(Dict)

	if err := json.Unmarshal(jsonDict, tmpDict); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON to Dict")
	}

	// Add KyuJitai to the dictionary
	tmpDict.appendKyujitai()

	return tmpDict, nil
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// appendKyujitai maps the Kyujitai (old kanjis) to the dictionary to speed up
// the search. Searching with the old kanjis will return the same result as the
// new kanjis.
func (dict *Dict) appendKyujitai() {
	// Add KyuJitai to the dictionary
	for _, tmpKanji := range *dict {
		if tmpKanji.KyuJitai == 0 {
			continue
		}

		// Convert the KyuJitai to a rune
		rKyuJitai := rune(tmpKanji.KyuJitai)

		// Add the KyuJitai to the dictionary
		if rKyuJitai != 0 {
			tmpKanji.IsKyuJitai = true

			(*dict)[rKyuJitai] = tmpKanji
		}
	}
}

// Find searches the given Kanji in the Joyo Kanji dictionary and returns the
// corresponding Kanji object.
// The returned boolean value indicates if the Kanji was found.
func (d Dict) Find(kanji rune) (Kanji, bool) {
	foundKanji, found := d[kanji]

	return foundKanji, found
}

// FixAsJoyo converts the given Kanji to the Joyo Kanji (regular use kanji) or
// Shin Jitai (new form of kanji) as much as possible.
//
// It will search the embedded Joyo Kanji dictionary then the Non-Joyo Kanji
// (old-new kanji mapping) dictionary.
//
// To add a new kanji, edit the file `non_joyo_old2new_map.go`.
func (d Dict) FixAsJoyo(kanji rune) rune {
	if !IsCJK(kanji) {
		return kanji
	}

	// Search the given Kanji in the Joyo Kanji dictionary
	tmpKanji, ok := d[kanji]
	if !ok {
		// Search the given Kanji in the Non-Joyo Kanji dictionary.
		// The returned kanji is the new kanji.
		if newKanji, ok := NonJoyoOld2NewMap[kanji]; ok {
			return newKanji
		}

		return kanji
	}

	if !tmpKanji.IsKyuJitai {
		return kanji
	}

	return rune(tmpKanji.ShinJitai)
}

// IsJoyoKanji returns true if the given Kanji is a Joyo Kanji.
func (d Dict) IsJoyoKanji(kanji rune) bool {
	if tmpKanji, ok := d[kanji]; ok {
		return rune(tmpKanji.ShinJitai) == kanji
	}

	return false
}

// IsKyuJitai returns true if the given Kanji is a KyuJitai (old kanji).
//
// Note that it only detects if the old kanji is registered in the dictionary.
func (d Dict) IsKyuJitai(kanji rune) bool {
	if tmpKanji, ok := d[kanji]; ok {
		return rune(tmpKanji.KyuJitai) == kanji
	}

	if _, ok := NonJoyoOld2NewMap[kanji]; ok {
		return true
	}

	return false
}

// KunYomi returns the KunYomi reading in hiragana of the given Kanji.
func (d Dict) KunYomi(kanji rune) []kana.Kanas {
	if tmpKanji, ok := d[kanji]; ok {
		return tmpKanji.Yomi.KunYomi
	}

	return nil
}

// LenJoyo counts and returns the number of Joyo Kanji elements registered in
// the dictionary.
//
// Notes:
//   - It does not count the KyuJitai (old kanji) entry.
//   - This method is not cached and is not suitable for frequent calls.
func (d Dict) LenJoyo() int {
	var count int

	for _, kanji := range d {
		if !kanji.IsKyuJitai {
			count++
		}
	}

	return count
}

// Marshal returns the JSON encoding of the dictionary.
//
// It is similar to MarshalJSON implementation but it strips the additional
// KyuJitai (old kanji) entries for search speed.
func (d Dict) Marshal() ([]byte, error) {
	byteJSON, err := json.Marshal(d.stripKyuJitai())

	return byteJSON, errors.Wrap(err, "failed to marshal the dictionary")
}

// MarshalIndent returns the indented JSON encoding of the dictionary.
func (d Dict) MarshalIndent(prefix string, indent string) ([]byte, error) {
	byteJSON, err := json.MarshalIndent(d.stripKyuJitai(), prefix, indent)

	return byteJSON, errors.Wrap(err, "failed to marshal the dictionary")
}

// OnYomi returns the OnYomi reading in katakana of the given Kanji.
func (d Dict) OnYomi(kanji rune) []kana.Kanas {
	if tmpKanji, ok := d[kanji]; ok {
		return tmpKanji.Yomi.OnYomi
	}

	return nil
}

// stripKyuJitai removes the elements only for speeding up the search from the
// Joyo Kanji dictionary. Elements such as ku_jitai as a key.
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
