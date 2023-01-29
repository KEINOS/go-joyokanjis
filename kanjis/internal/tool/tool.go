package tool

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"

	"github.com/pkg/errors"
)

func ExtractGzipGobToDict(src io.Reader, dest any) error {
	gzReader, err := gzip.NewReader(src)
	if err != nil {
		return errors.Wrap(err, "failed to create a gzip reader")
	}

	defer gzReader.Close()

	// Extract gzipped data
	var gobData []byte

	buf := bytes.NewBuffer(gobData)
	if _, err := io.Copy(buf, gzReader); err != nil {
		return errors.Wrap(err, "failed to extract the GZipped Gob encoded data")
	}

	// Decode the gob data to kanji.Dict object and set to kanjiDict.
	rawGob := bytes.NewBuffer(buf.Bytes())
	if err = gob.NewDecoder(rawGob).Decode(dest); err != nil {
		return errors.Wrap(err, "failed to decode the Gob encoded data")
	}

	return nil
}
