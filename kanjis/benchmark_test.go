package kanjis

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mackerelio/go-osstat/memory"
)

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

// memoryGet is a copy of memory.Get function to ease testing.
var MemoryGet = memory.Get

// AvailableMemory returns the amount of current available free memory.
//
// It will error if it fails to get the memory information. Mostly on platforms
// such as NetBSD and OpenBSD.
func AvailableMemory(b *testing.B) uint64 {
	b.Helper()

	mem, err := MemoryGet()
	if err != nil {
		b.Fatal(err)
	}

	return mem.Free
}
