[![go1.19+](https://img.shields.io/badge/Go-1.19+-blue?logo=go)](https://github.com/KEINOS/go-joyokanjis/blob/main/go.mod "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-gisty.svg)](https://pkg.go.dev/github.com/KEINOS/go-joyokanjis/kanjis)

# go-joyokanjis

`go-joyokanjis` is a simple Go library for Japanese writings that lint or determines whether a given kanji character is a [jōyō-kanji](https://en.wikipedia.org/wiki/J%C5%8Dy%C5%8D_kanji) (常用漢字, regular-use Chinese characters in Japan) or not.

Useful for uniforming the kanji characters in the writings. Currently, it supports 2,136 jōyō-kanji characters.

> __Note__: This library is based on Unicode (UTF-8) and does not support other Japanese character encodings such as JIS/SJIS/EUC/etc.

## Usage

```go
go get "github.com/KEINOS/go-joyokanjis"
```

```go
import "github.com/KEINOS/go-joyokanjis/kanjis"
```

## Examples

### Detection

```go
// Detect if a given kanji is Joyo Kanji or not.
func ExampleIsJoyokanji() {
    newKanji := '漢'
    if kanjis.IsJoyokanji(newKanji) {
        fmt.Printf("%s (0x%x) is Joyo Kanji\n", string(newKanji), newKanji)
    }

    oldKanji := '漢'
    if !kanjis.IsJoyokanji(oldKanji) {
        fmt.Printf("%s (0x%x) is not a Joyo Kanji\n", string(oldKanji), oldKanji)
    }

    // Output:
    // 漢 (0x6f22) is Joyo Kanji
    // 漢 (0xfa47) is not a Joyo Kanji
}
```

### Fixing

```go
// Fix a string to replace all old kanji characters with Joyo Kanji (only if the
// old kanji is assigned to Joyo Kanji).
//
// Suitable if the input is less than 320 Bytes.
func ExampleFixStringAsJoyo() {
    input := "これは舊漢字です。And this is not a kanji."
    output := kanjis.FixStringAsJoyo(input)

    fmt.Println(output)
    // Output: これは旧漢字です。And this is not a kanji.
}
```

```go
// Fix a file to replace all old kanji characters with Joyo Kanji (only if the
// old kanji is assigned to Joyo Kanji).
//
// This function is suitable if the input is larger than 320 Bytes.
func ExampleFixFileAsJoyo() {
    // File content
    input := strings.NewReader(heredoc.Doc(`
        いざ、これより樂しまむ、
        仕置を受くる憂なく、
        遊びたのしむ時ぞ來ぬ、
        時ぞ來ぬれば、いちはやく、
        讀本などは投げ捨てて行く。
        ――學校休暇の歌`))

    // Output buffer
    var output bytes.Buffer

    // Parse and fix to Joyo Kanji
    if err := kanjis.FixFileAsJoyo(input, &output); err != nil {
        log.Fatal(err)
    }

    fmt.Println(output.String())
    // Output:
    // いざ、これより楽しまむ、
    // 仕置を受くる憂なく、
    // 遊びたのしむ時ぞ来ぬ、
    // 時ぞ来ぬれば、いちはやく、
    // 読本などは投げ捨てて行く。
    // ――学校休暇の歌
}
```

```go
func ExampleIgnore() {
    const input = "私は渡邉です。"

    {
        // Add '邉' and '邊' to be ignored when fixing.
        kanjis.Ignore('邉', '邊')

        fmt.Println("Fix with Ignore:", kanjis.FixStringAsJoyo(input))
    }
    {
        // Clear the ignore list.
        kanjis.ResetIgnore()

        fmt.Println("Fix with no-ignore:", kanjis.FixStringAsJoyo(input))
    }
    // Output:
    // Fix with Ignore: 私は渡邉です。
    // Fix with no-ignore: 私は渡辺です。
}
```

## Benchmark

```text
goos: darwin
goarch: amd64
pkg: github.com/KEINOS/go-joyokanjis/kanjis
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
                              │ bench_results.txt │
                              │      sec/op       │
_small_size/FixStringAsJoyo-4        384.7n ±  3%
_small_size/FixFileAsJoyo-4          5.935µ ± 89%
_big_size/FixStringAsJoyo-4          308.5µ ±  2%
_big_size/FixFileAsJoyo-4            4.705µ ± 15%
geomean                              7.587µ

                              │ bench_results.txt │
                              │       B/op        │
_small_size/FixStringAsJoyo-4          32.00 ± 0%
_small_size/FixFileAsJoyo-4          8.156Ki ± 0%
_big_size/FixStringAsJoyo-4          46.62Ki ± 0%
_big_size/FixFileAsJoyo-4            8.156Ki ± 0%
geomean                              3.138Ki

                              │ bench_results.txt │
                              │     allocs/op     │
_small_size/FixStringAsJoyo-4          1.000 ± 0%
_small_size/FixFileAsJoyo-4            4.000 ± 0%
_big_size/FixStringAsJoyo-4            2.000 ± 0%
_big_size/FixFileAsJoyo-4              4.000 ± 0%
geomean                                2.378
```
