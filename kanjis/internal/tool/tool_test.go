package tool

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	MsgError   string
	Name       string
	ForceError bool
}

func (t TestStruct) Read(p []byte) (n int, err error) {
	if t.ForceError {
		return 0, errors.New(t.MsgError)
	}

	return 0, io.EOF
}

func generateGzippedGobData(val string) (r io.Reader, e error) {
	var gZippedGobEncData bytes.Buffer

	data := TestStruct{
		Name: val,
	}

	// Encode to gob format.
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(&data); err != nil {
		return nil, errors.Wrap(err, "failed to Gob encode the data")
	}

	// Gzip the gob data.
	gzWriter := gzip.NewWriter(&gZippedGobEncData)
	if _, err := gzWriter.Write(buf.Bytes()); err != nil {
		return nil, errors.Wrap(err, "failed to GZip the Gob encoded data")
	}

	defer func() {
		e = gzWriter.Close()
	}()

	return &gZippedGobEncData, nil
}

func TestExtractGzipGobToDict_golden(t *testing.T) {
	// Generate a golden data.
	expectValue := t.Name()

	archivedData, err := generateGzippedGobData(expectValue)
	require.NoError(t, err, "failed to create test data (archive of gob encoded data)")

	var testObj TestStruct

	// Test
	err = ExtractGzipGobToDict(archivedData, &testObj)
	require.NoError(t, err, "failed to extract the GZipped Gob encoded data")

	actualValue := testObj.Name

	require.Equal(t, expectValue, actualValue,
		"the extracted object does not contain the expected value")
}

func TestExtractGzipGobToDict_read_fail(t *testing.T) {
	dummyObj := &TestStruct{
		Name:       "test data",
		ForceError: true,
		MsgError:   "forced error",
	}

	var testObj TestStruct

	// Test
	err := ExtractGzipGobToDict(dummyObj, &testObj)

	require.Error(t, err,
		"it should fail if the given reader fails to read")
	require.Empty(t, testObj,
		"the dest object should be empty on error")
	require.Contains(t, err.Error(), "failed to create a gzip reader",
		"it should contain the error reason")
	require.Contains(t, err.Error(), "forced error",
		"it should contain the wrapped error reason")
}

func TestExtractGzipGobToDict_fail_copy(t *testing.T) {
	// Generate a golden data.
	archivedData, err := generateGzippedGobData(t.Name())
	require.NoError(t, err, "failed to create test data (archive of gob encoded data)")

	// Backup and defer restore the gzipNewReader function.
	oldIoCopy := ioCopy
	defer func() {
		ioCopy = oldIoCopy
	}()

	// Mock the gzipNewReader function to return an empty reader.
	ioCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("forced werrors")
	}

	var testObj TestStruct

	// Test
	err = ExtractGzipGobToDict(archivedData, &testObj)

	require.Error(t, err,
		"it should fail if the io.Copy fails to read/write")
	require.Contains(t, err.Error(), "failed to extract the GZipped Gob encoded data",
		"it should contain the error reason")
}
