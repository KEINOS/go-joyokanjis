package kanji

import "github.com/goccy/go-json"

// Unmarshal parses the Joyo Kanji dictionary in JSON format to the kanji.Dict
// type.
//
// If the registered joyo kanji has an old kanji (kyu jitai), an alias key will
// be added to the dictionary to speed up the search.
//
// See the following URL for the format of the JSON byte array:
//
//	https://gist.github.com/KEINOS/fb660943484008b7f5297bb627e0e1b1#format
func Unmarshal(jsonDict []byte) (Dict, error) {
	var tmpDict Dict

	if err := json.Unmarshal(jsonDict, &tmpDict); err != nil {
		return nil, err
	}

	// Add KyuJitai to the dictionary
	appendKyujitai(&tmpDict)

	return tmpDict, nil
}

// appendKyujitai maps the Kyujitai (old kanjis) to the dictionary to speed up
// the search. Searching with the old kanjis will return the same result as the
// new kanjis.
func appendKyujitai(dict *Dict) {
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
