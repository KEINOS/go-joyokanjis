# go-joyokanjis

`go-joyokanjis` is a simple Go library for Japanese writings that determines whether a given kanji character is [jōyō-kanji](https://en.wikipedia.org/wiki/J%C5%8Dy%C5%8D_kanji) (常用漢字, regular-use Chinese characters in Japan) or not.

This library is based on Unicode and does not support other Japanese character encodings such as JIS/SJIS/EUC/etc.

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
// Fix a file to replace all ald kanji characters with Joyo Kanji (only if the
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
