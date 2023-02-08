/*
This package generates the Joyo-kanji dictionary to be embedded in the package.

It will download the dictionary in JSON and converts to a gob encoded format,
then gzips it to be embedded in the package.

To run/generate, use the following command from the root of the project:

	go generate ./...
*/
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/KEINOS/go-joyokanjis/kanjis/kanji"
	"github.com/pkg/errors"
)

// urlDictSource is the URL to the JSON dictionary.
const (
	urlDictSourceDefault = "https://gist.githubusercontent.com/KEINOS/fb660943484008b7f5297bb627e0e1b1/raw/joyo2010.json"
	levelCompressDefault = gzip.BestCompression
)

var (
	urlDictSource  string
	pathJSONInput  string
	pathGobOutput  string
	pathGzipOutput string

	levelCompress = levelCompressDefault
)

func init() {
	urlDictSource = urlDictSourceDefault
	pathJSONInput = filepath.Join("internal", "json", "joyo2010.json")
	pathGobOutput = filepath.Join("internal", "gob", "dict.gob")
	pathGzipOutput = filepath.Join("internal", "gzgob", "dict.gzip")
}

func main() {
	// Read/download the JSON dictionary file.
	if !fileExists(pathJSONInput) {
		fmt.Println("Downloading JSON file...")
		exitOnError(downloadDictJSON(pathJSONInput))
	}

	dataJSON, err := os.ReadFile(pathJSONInput)
	exitOnError(err)

	// Parse JSON to kanji.Dict.
	dict, err := kanji.NewDict(dataJSON)
	exitOnError(err)

	// Convert the dictionary object to a gob encoded format.
	buf := bytes.NewBuffer(nil)
	exitOnError(gob.NewEncoder(buf).Encode(&dict))

	// Save the dictionary to a gob file.
	exitOnError(os.WriteFile(pathGobOutput, buf.Bytes(), os.ModePerm))

	// Compress the gob file to a gz file.
	ptrOut, err := os.Create(pathGzipOutput)
	exitOnError(err)

	gw, err := gzip.NewWriterLevel(ptrOut, levelCompress)
	exitOnError(err)

	defer gw.Close()

	// Write the compressed data to the file.
	src, err := os.Open(pathGobOutput)
	exitOnError(err)

	defer src.Close()

	_, err = io.Copy(gw, src)
	exitOnError(err)

	fmt.Println("OK")
}

func downloadDictJSON(to string) error {
	resp, err := http.Get(urlDictSource)
	if err != nil {
		return errors.Wrap(err, "failed to download the JSON dictionary")
	}

	defer resp.Body.Close()

	out, err := os.Create(to)
	if err != nil {
		return errors.Wrap(err, "failed to create a file to save the JSON dictionary")
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return errors.Wrap(err, "failed to copy the downloaded data to the target file")
}

// exitOnError exits the progrom if err is not nil. It will panic to let defer
// functions run.
func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// fileExists returns true if the given path is an existing file. Returns false
// if the path is a directory or does not exist.
func fileExists(pathFile string) bool {
	info, err := os.Stat(pathFile)
	if err == nil {
		return !info.IsDir()
	}

	return false
}
