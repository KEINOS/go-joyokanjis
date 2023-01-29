package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Code-Hex/dd"
	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
)

func Example() {
	pathFileGob := filepath.Join("gob", "dict.gob")

	rawGob, err := os.ReadFile(pathFileGob)
	if err != nil {
		log.Fatal(err)
	}

	var loadedDict kanji.Dict

	buf := bytes.NewBuffer(rawGob)
	if err = gob.NewDecoder(buf).Decode(&loadedDict); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dd.Dump(loadedDict['亞']))
	// Output:
	// kanji.Kanji{
	//   Yomi: kanji.Yomi{
	//     OnYomi: []string{
	//       "ア",
	//     },
	//     KunYomi: ([]string)(nil),
	//     ExampleYomi: ([]string)(nil),
	//   },
	//   ShinJitai: 20124,
	//   KyuJitai: 20126,
	//   IsKyuJitai: true,
	// }
}
