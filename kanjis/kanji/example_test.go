package kanji_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
)

func Example() {
	// Sample JSON data. The old kanji "滯"(28399) is mapped to the new kanji "滞".
	// For the JSON format see: https://gist.github.com/KEINOS/fb660943484008b7f5297bb627e0e1b1#format
	sampleJSON := `{
		"28382": {
			"joyo_kanji": "滞",
			"kyu_jitai": "滯",
			"yomi": {
				"on_yomi": ["タイ"],
				"kun_yomi": ["とどこお"],
				"example_yomi": ["とどこお-る"]
			},
			"raw_info": "滞\t滯\t13\t7S\t\tタイ、とどこお-る"
		}
	}`

	// Unmarshal JSON to Dict
	tmpDict, err := kanji.Unmarshal([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// 滞 = 28382, new kanji
	fmt.Println("滞 -->", tmpDict.IsJoyokanji('滞'), tmpDict.OnYomi('滞'), tmpDict.KunYomi('滞'))
	// Since "滞" contains an old kanji "滯"(28399), a new entry with the old kanji
	// is also mapped to the dictionary.
	fmt.Println("滯 -->", tmpDict.IsJoyokanji('滯'), tmpDict.OnYomi('滯'), tmpDict.KunYomi('滯'))

	// Print the structure of the Kanji object (not in JSON format)
	fmt.Println(tmpDict['滞'])
	// Output:
	// 滞 --> true [タイ] [とどこお]
	// 滯 --> false [タイ] [とどこお]
	// kanji.Kanji{
	//   Yomi: kanji.Yomi{
	//     OnYomi: []string{
	//       "タイ",
	//     },
	//     KunYomi: []string{
	//       "とどこお",
	//     },
	//     ExampleYomi: []string{
	//       "とどこお-る",
	//     },
	//   },
	//   ShinJitai: 28382,
	//   KyuJitai: 28399,
	//   IsKyuJitai: false,
	// }
}

func ExampleDict_Marshal() {
	sampleJSON := `{
		"28382": {
			"joyo_kanji": "滞",
			"kyu_jitai": "滯",
			"yomi": {
				"on_yomi": ["タイ"],
				"kun_yomi": ["とどこお"],
				"example_yomi": ["とどこお-る"]
			},
			"raw_info": "滞\t滯\t13\t7S\t\tタイ、とどこお-る"
		}
	}`

	// Unmarshal JSON to Dict. It will add kyujitai entry to the Dict to speed
	// up the search.
	tmpDict, err := kanji.Unmarshal([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Marshal the Dict obj to JSON. It trimms the kyujitai entry for serchability.
	// To get the full JSON, use the json.Marshal() function instead.
	// To pretty print the JSON, use the Dict.MarshalIndent() method instead.
	jsonDict, err := tmpDict.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDict))
	// Output:
	// {"28382":{"yomi":{"on_yomi":["タイ"],"kun_yomi":["とどこお"],"example_yomi":["とどこお-る"]},"joyo_kanji":"滞","kyu_jitai":"滯"}}
}
