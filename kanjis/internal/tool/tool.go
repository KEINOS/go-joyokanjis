package tool

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

// io.opy is a copy of io.Copy() to ease testing.
var ioCopy = io.Copy

// ExtractGzipGobToDict extracts the GZipped Gob encoded data from the source
// and decodes it to the destination.
//
// Note that the dest argument accepts any type, but is assumed to be of type
// kanji.Dict.
func ExtractGzipGobToDict(src io.Reader, dest any) error {
	gzReader, err := gzip.NewReader(src)
	if err != nil {
		return errors.Wrap(err, "failed to create a gzip reader")
	}

	defer gzReader.Close()

	// Extract gzipped data
	var gobData []byte

	buf := bytes.NewBuffer(gobData)
	if _, err := ioCopy(buf, gzReader); err != nil {
		return errors.Wrap(err, "failed to extract the GZipped Gob encoded data")
	}

	// Decode the gob data to kanji.Dict object and set to kanjiDict.
	rawGob := bytes.NewBuffer(buf.Bytes())

	err = gob.NewDecoder(rawGob).Decode(dest)

	return errors.Wrap(err, "failed to decode the Gob encoded data")
}
