package kanjis

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ----------------------------------------------------------------------------
//  Benchmark
// ----------------------------------------------------------------------------

func Benchmark_small_size(b *testing.B) {
	const input = "これは舊漢字です。"

	b.Run("FixStringAsJoyo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FixStringAsJoyo(input)
		}
	})

	b.Run("FixFileAsJoyo", func(b *testing.B) {
		ptrInput := strings.NewReader(input)
		var output bytes.Buffer

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FixFileAsJoyo(ptrInput, &output)
		}
	})
}

func Benchmark_big_size(b *testing.B) {
	input := getData(b)

	b.Run("FixStringAsJoyo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FixStringAsJoyo(input)
		}
	})

	b.Run("FixFileAsJoyo", func(b *testing.B) {
		ptrInput := strings.NewReader(input)
		var output bytes.Buffer

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FixFileAsJoyo(ptrInput, &output)
		}
	})
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

var testDataBigSize string

func getData(b *testing.B) string {
	b.Helper()

	if len(testDataBigSize) > 0 {
		return testDataBigSize
	}

	pathFileData := filepath.Join("testdata", "ekiden_basha.txt")
	data, err := os.ReadFile(pathFileData)
	if err != nil {
		b.Fatalf("ERROR: %s", err)
	}

	testDataBigSize = string(data)

	return testDataBigSize
}
