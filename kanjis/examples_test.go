package kanjis_test

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/KEINOS/go-joyokanjis/kanjis"
	"github.com/MakeNowJust/heredoc"
)

func Example() {
	input := heredoc.Doc(`
		いざ、これより樂しまむ、
		仕置を受くる憂なく、
		遊びたのしむ時ぞ來ぬ、
		時ぞ來ぬれば、いちはやく、
		讀本などは投げ捨てて行く。
		――學校休暇の歌`)

	output := kanjis.FixStringAsJoyo(input)

	fmt.Println(output)
	// Output:
	// いざ、これより楽しまむ、
	// 仕置を受くる憂なく、
	// 遊びたのしむ時ぞ来ぬ、
	// 時ぞ来ぬれば、いちはやく、
	// 読本などは投げ捨てて行く。
	// ――学校休暇の歌
}

func ExampleFixRuneAsJoyo() {
	for _, test := range []struct {
		input  rune
		expect rune
	}{
		{input: '漢', expect: '漢'},
		{input: '漢', expect: '漢'},
		{input: '巣', expect: '巣'},
		{input: '巢', expect: '巣'},
		{input: 'a', expect: 'a'},
		{input: 'あ', expect: 'あ'},
		{input: 'ア', expect: 'ア'},
	} {
		expect := test.expect
		actual := kanjis.FixRuneAsJoyo(test.input)

		if expect != actual {
			log.Fatalf("ERROR: Expected %q but got %q", string(expect), string(actual))
		}
	}

	fmt.Println("OK")
	// Output: OK
}

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

func ExampleFixStringAsJoyo() {
	input := "これは舊漢字です。"
	output := kanjis.FixStringAsJoyo(input)

	fmt.Println(output)
	// Output: これは旧漢字です。
}

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

func ExampleIsJoyoKanji() {
	newKanji := '漢'
	if kanjis.IsJoyoKanji(newKanji) {
		fmt.Printf("%s (0x%x) is Joyo Kanji\n", string(newKanji), newKanji)
	}

	oldKanji := '漢'
	if !kanjis.IsJoyoKanji(oldKanji) {
		fmt.Printf("%s (0x%x) is not a Joyo Kanji\n", string(oldKanji), oldKanji)
	}

	// Output:
	// 漢 (0x6f22) is Joyo Kanji
	// 漢 (0xfa47) is not a Joyo Kanji
}

func ExampleIsKyuJitai() {
	for index, test := range []struct {
		input  rune
		expect bool
	}{
		{input: '漢', expect: false}, // New kanji
		{input: '漢', expect: true},  // Old kanji
		{input: '亙', expect: true},  // Old kanji but not in Joyo Kanji list
		{input: 'a', expect: false}, // Not a kanji
	} {
		expect := test.expect
		actual := kanjis.IsKyuJitai(test.input)

		if expect != actual {
			log.Fatalf("test #%d failed: IsKyuJitai('%s') expected to be %t but got %t",
				index, string(test.input), expect, actual)
		}
	}

	fmt.Println("OK")
	// Output: OK
}

func ExampleLenDict() {
	fmt.Println(kanjis.LenDict())
	// Output: 2136
}
