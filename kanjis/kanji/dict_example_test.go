package kanji_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
)

// ============================================================================
//  Dict type
// ============================================================================

// ----------------------------------------------------------------------------
//  Dict.Find()
// ----------------------------------------------------------------------------

func ExampleDict_Find() {
	// Sample JSON dictionary.
	// In this example only one kanji (new and old) is registered.
	sampleJSON := `{
		"27005": {
			"joyo_kanji": "楽",
			"kyu_jitai": "樂",
			"yomi": {
				"on_yomi": [
					"ガク",
					"ラク"
				],
				"kun_yomi": [
					"たの"
				],
				"example_yomi": [
					"たの-しい",
					"たの-しむ"
				]
			},
			"raw_info": "楽\t樂\t13\t2\t\tガク、ラク、たの-しい、たの-しむ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	if foundKanji, ok := tmpDict.Find('楽'); ok {
		fmt.Println("Is kyu-jitai?:", foundKanji.IsKyuJitai)
		fmt.Println("Shin jitai (new form kanji):", foundKanji.ShinJitai)
		fmt.Println("Kyu jitai (old form kanji):", foundKanji.KyuJitai)
		fmt.Println("On-reading:", foundKanji.Yomi.OnYomi)
		fmt.Println("Kun-reading:", foundKanji.Yomi.KunYomi)
		fmt.Println("Examples of the readings:", foundKanji.Yomi.ExampleYomi)
	}
	// Output:
	// Is kyu-jitai?: false
	// Shin jitai (new form kanji): 楽
	// Kyu jitai (old form kanji): 樂
	// On-reading: [ガク ラク]
	// Kun-reading: [たの]
	// Examples of the readings: [たの-しい たの-しむ]
}

// ----------------------------------------------------------------------------
//  Dict.FixAsJoyo()
// ----------------------------------------------------------------------------

func ExampleDict_FixAsJoyo() {
	// Sample JSON dictionary.
	// In this example only one kanji (new and old) is registered.
	sampleJSON := `{
		"27005": {
			"joyo_kanji": "楽",
			"kyu_jitai": "樂",
			"yomi": {
				"on_yomi": [
					"ガク",
					"ラク"
				],
				"kun_yomi": [
					"たの"
				],
				"example_yomi": [
					"たの-しい",
					"たの-しむ"
				]
			},
			"raw_info": "楽\t樂\t13\t2\t\tガク、ラク、たの-しい、たの-しむ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Fix if rune is a registered old kanji to the new kanji.
	for index, test := range []struct {
		input  rune
		expect rune
	}{
		{input: '樂', expect: '楽'}, // Registered old kanji
		{input: '楽', expect: '楽'}, // Registered new kanji
		{input: 'あ', expect: 'あ'}, // Non kanji (not CJK)
		{input: '滯', expect: '滯'}, // Non registered kanji in the JSON dictionary
		{input: '亙', expect: '亘'}, // Non registered but NonJoyoOld2NewMap has the mapping
	} {
		expect := test.expect
		actual := tmpDict.FixAsJoyo(test.input)

		fmt.Printf("#%d: %s expected to be: %s, got: %s\n",
			index+1, string(test.input), string(expect), string(actual))
	}
	// Output:
	// #1: 樂 expected to be: 楽, got: 楽
	// #2: 楽 expected to be: 楽, got: 楽
	// #3: あ expected to be: あ, got: あ
	// #4: 滯 expected to be: 滯, got: 滯
	// #5: 亙 expected to be: 亘, got: 亘
}

// ----------------------------------------------------------------------------
//  Dict.IsJoyoKanji()
// ----------------------------------------------------------------------------

func ExampleDict_IsJoyoKanji() {
	// Sample JSON dictionary.
	// In this example only one kanji (new and old) is registered.
	sampleJSON := `{
		"27005": {
			"joyo_kanji": "楽",
			"kyu_jitai": "樂",
			"yomi": {
				"on_yomi": [
					"ガク",
					"ラク"
				],
				"kun_yomi": [
					"たの"
				],
				"example_yomi": [
					"たの-しい",
					"たの-しむ"
				]
			},
			"raw_info": "楽\t樂\t13\t2\t\tガク、ラク、たの-しい、たの-しむ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Search the registered kanji
	fmt.Println("Is '楽' a joyo kanji?:", tmpDict.IsJoyoKanji('楽'))
	fmt.Println("Is '樂' a joyo kanji?:", tmpDict.IsJoyoKanji('樂'))
	// Non-registered kanji/rune returns false.
	fmt.Println("Is '忍' a joyo kanji?:", tmpDict.IsJoyoKanji('忍'))
	fmt.Println("Is '아' a joyo kanji?:", tmpDict.IsJoyoKanji('아'))
	// Output:
	// Is '楽' a joyo kanji?: true
	// Is '樂' a joyo kanji?: false
	// Is '忍' a joyo kanji?: false
	// Is '아' a joyo kanji?: false
}

// ----------------------------------------------------------------------------
//  Dict.IsKyuJitai()
// ----------------------------------------------------------------------------

func ExampleDict_IsKyuJitai() {
	// Sample JSON dictionary.
	// In this example only one kanji (new and old) is registered.
	sampleJSON := `{
		"27005": {
			"joyo_kanji": "楽",
			"kyu_jitai": "樂",
			"yomi": {
				"on_yomi": [
					"ガク",
					"ラク"
				],
				"kun_yomi": [
					"たの"
				],
				"example_yomi": [
					"たの-しい",
					"たの-しむ"
				]
			},
			"raw_info": "楽\t樂\t13\t2\t\tガク、ラク、たの-しい、たの-しむ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Search the registered kanji
	fmt.Println("Is '楽' a kyu-jitai?:", tmpDict.IsKyuJitai('楽'))
	fmt.Println("Is '樂' a kyu-jitai?:", tmpDict.IsKyuJitai('樂'))
	// Non-registered kanji/rune returns false.
	fmt.Println("Is '忍' a kyu-jitai?:", tmpDict.IsKyuJitai('忍'))
	fmt.Println("Is '아' a kyu-jitai?:", tmpDict.IsJoyoKanji('아'))
	// Output:
	// Is '楽' a kyu-jitai?: false
	// Is '樂' a kyu-jitai?: true
	// Is '忍' a kyu-jitai?: false
	// Is '아' a kyu-jitai?: false
}

// ----------------------------------------------------------------------------
//  Dict.KunYomi()
// ----------------------------------------------------------------------------

func ExampleDict_KunYomi() {
	// Sample JSON dictionary.
	sampleJSON := `{
		"32905": {
			"joyo_kanji": "肉",
			"yomi": {
				"on_yomi": [
					"ニク"
				]
			},
			"raw_info": "肉\t\t6\t2\t\tニク"
		},
		"32908": {
			"joyo_kanji": "肌",
			"yomi": {
				"kun_yomi": [
					"はだ"
				]
			},
			"raw_info": "肌\t\t6\t7S\t1981\tはだ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("肉:", tmpDict.KunYomi('肉'))
	fmt.Println("肌:", tmpDict.KunYomi('肌'))
	fmt.Println("忍:", tmpDict.KunYomi('忍')) // Not in the current dictionary
	// Output:
	// 肉: []
	// 肌: [はだ]
	// 忍: []
}

// ----------------------------------------------------------------------------
//  Dict.LenJoyo
// ----------------------------------------------------------------------------

func ExampleDict_LenJoyo() {
	// Sample JSON dictionary.
	// In this example only one kanji (with kyu-jitai, old kanji) is registered.
	sampleJSON := `{
		"27005": {
			"joyo_kanji": "楽",
			"kyu_jitai": "樂",
			"yomi": {
				"on_yomi": [
					"ガク",
					"ラク"
				],
				"kun_yomi": [
					"たの"
				],
				"example_yomi": [
					"たの-しい",
					"たの-しむ"
				]
			},
			"raw_info": "楽\t樂\t13\t2\t\tガク、ラク、たの-しい、たの-しむ"
		}
	}`

	// Create a new dictionary from the JSON dictionary. Since the registered
	// kanji has a kyu-jitai, it will add an additional element to the dictionary
	// to speed up the search.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Expect 1
	fmt.Println("Total joyo kanji in the dictionary:", tmpDict.LenJoyo())
	// Expect 2 (1 joyo kanji + 1 kyu-jitai)
	fmt.Println("Total elements in the dictionary:", len(*tmpDict))
	// Output:
	// Total joyo kanji in the dictionary: 1
	// Total elements in the dictionary: 2
}

// ----------------------------------------------------------------------------
//  Dict.Marshal()
// ----------------------------------------------------------------------------

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

	// Create the dictionary (unmarshal JSON to Dict). It will add kyujitai
	// entry to the Dict to speed up the search.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// Convert/marshal the Dict obj to JSON. It trimms the additional kyujitai
	// entry for serchability. To pretty print the JSON, use the Dict.MarshalIndent()
	// method instead.
	jsonDictTrimmed, err := tmpDict.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDictTrimmed))

	// To get the full entries (including the additional entries for searching),
	// use the ordinary json.Marshal() function.
	jsonDictRaw, err := json.Marshal(tmpDict)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDictRaw))
	// Output:
	// {"28382":{"yomi":{"on_yomi":["タイ"],"kun_yomi":["とどこお"],"example_yomi":["とどこお-る"]},"joyo_kanji":"滞","kyu_jitai":"滯"}}
	// {"28382":{"yomi":{"on_yomi":["タイ"],"kun_yomi":["とどこお"],"example_yomi":["とどこお-る"]},"joyo_kanji":"滞","kyu_jitai":"滯"},"28399":{"yomi":{"on_yomi":["タイ"],"kun_yomi":["とどこお"],"example_yomi":["とどこお-る"]},"joyo_kanji":"滞","kyu_jitai":"滯"}}
}

// ----------------------------------------------------------------------------
// Dict.MarshalIndent()
// ----------------------------------------------------------------------------

func ExampleDict_MarshalIndent() {
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

	// Create the dictionary (unmarshal JSON to Dict). It will add kyujitai entry
	// to the Dict to speed up the search.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	// It is similar to Dict.Marshal() but it pretty prints the JSON.
	//
	// Note that additional entries for searching are trimmed. For printing
	// the full entries, use the ordinary json.MarshalIndent() function.
	prefix := ""
	indent := "  "
	jsonDictTrimmed, err := tmpDict.MarshalIndent(prefix, indent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDictTrimmed))
	// Output:
	// {
	//   "28382": {
	//     "yomi": {
	//       "on_yomi": [
	//         "タイ"
	//       ],
	//       "kun_yomi": [
	//         "とどこお"
	//       ],
	//       "example_yomi": [
	//         "とどこお-る"
	//       ]
	//     },
	//     "joyo_kanji": "滞",
	//     "kyu_jitai": "滯"
	//   }
	// }
}

// ----------------------------------------------------------------------------
//  Dict.OnYomi()
// ----------------------------------------------------------------------------

func ExampleDict_OnYomi() {
	// Sample JSON dictionary.
	sampleJSON := `{
		"32905": {
			"joyo_kanji": "肉",
			"yomi": {
				"on_yomi": [
					"ニク"
				]
			},
			"raw_info": "肉\t\t6\t2\t\tニク"
		},
		"32908": {
			"joyo_kanji": "肌",
			"yomi": {
				"kun_yomi": [
					"はだ"
				]
			},
			"raw_info": "肌\t\t6\t7S\t1981\tはだ"
		}
	}`

	// Create a new dictionary from the JSON dictionary.
	tmpDict, err := kanji.NewDict([]byte(sampleJSON))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("肉:", tmpDict.OnYomi('肉'))
	fmt.Println("肌:", tmpDict.OnYomi('肌'))
	fmt.Println("忍:", tmpDict.OnYomi('忍')) // Not in the current dictionary
	// Output:
	// 肉: [ニク]
	// 肌: []
	// 忍: []
}
