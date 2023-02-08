package kana_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KEINOS/go-joyokanjis/kanjis/kana"
	"github.com/MakeNowJust/heredoc"
)

func ExampleToKatakana() {
	for i := 'ぁ'; i < 'ゖ'+1; i++ {
		k := kana.ToKatakana(i)

		fmt.Println(string(i), "-->", string(k))
	}
	// Output:
	// ぁ --> ァ
	// あ --> ア
	// ぃ --> ィ
	// い --> イ
	// ぅ --> ゥ
	// う --> ウ
	// ぇ --> ェ
	// え --> エ
	// ぉ --> ォ
	// お --> オ
	// か --> カ
	// が --> ガ
	// き --> キ
	// ぎ --> ギ
	// く --> ク
	// ぐ --> グ
	// け --> ケ
	// げ --> ゲ
	// こ --> コ
	// ご --> ゴ
	// さ --> サ
	// ざ --> ザ
	// し --> シ
	// じ --> ジ
	// す --> ス
	// ず --> ズ
	// せ --> セ
	// ぜ --> ゼ
	// そ --> ソ
	// ぞ --> ゾ
	// た --> タ
	// だ --> ダ
	// ち --> チ
	// ぢ --> ヂ
	// っ --> ッ
	// つ --> ツ
	// づ --> ヅ
	// て --> テ
	// で --> デ
	// と --> ト
	// ど --> ド
	// な --> ナ
	// に --> ニ
	// ぬ --> ヌ
	// ね --> ネ
	// の --> ノ
	// は --> ハ
	// ば --> バ
	// ぱ --> パ
	// ひ --> ヒ
	// び --> ビ
	// ぴ --> ピ
	// ふ --> フ
	// ぶ --> ブ
	// ぷ --> プ
	// へ --> ヘ
	// べ --> ベ
	// ぺ --> ペ
	// ほ --> ホ
	// ぼ --> ボ
	// ぽ --> ポ
	// ま --> マ
	// み --> ミ
	// む --> ム
	// め --> メ
	// も --> モ
	// ゃ --> ャ
	// や --> ヤ
	// ゅ --> ュ
	// ゆ --> ユ
	// ょ --> ョ
	// よ --> ヨ
	// ら --> ラ
	// り --> リ
	// る --> ル
	// れ --> レ
	// ろ --> ロ
	// ゎ --> ヮ
	// わ --> ワ
	// ゐ --> ヰ
	// ゑ --> ヱ
	// を --> ヲ
	// ん --> ン
	// ゔ --> ヴ
	// ゕ --> ヵ
	// ゖ --> ヶ
}

func ExampleToHiragana() {
	for i := 'ァ'; i < 'ヶ'+1; i++ {
		k := kana.ToHiragana(i)

		fmt.Println(string(i), "-->", string(k))
	}
	// Output:
	// ァ --> ぁ
	// ア --> あ
	// ィ --> ぃ
	// イ --> い
	// ゥ --> ぅ
	// ウ --> う
	// ェ --> ぇ
	// エ --> え
	// ォ --> ぉ
	// オ --> お
	// カ --> か
	// ガ --> が
	// キ --> き
	// ギ --> ぎ
	// ク --> く
	// グ --> ぐ
	// ケ --> け
	// ゲ --> げ
	// コ --> こ
	// ゴ --> ご
	// サ --> さ
	// ザ --> ざ
	// シ --> し
	// ジ --> じ
	// ス --> す
	// ズ --> ず
	// セ --> せ
	// ゼ --> ぜ
	// ソ --> そ
	// ゾ --> ぞ
	// タ --> た
	// ダ --> だ
	// チ --> ち
	// ヂ --> ぢ
	// ッ --> っ
	// ツ --> つ
	// ヅ --> づ
	// テ --> て
	// デ --> で
	// ト --> と
	// ド --> ど
	// ナ --> な
	// ニ --> に
	// ヌ --> ぬ
	// ネ --> ね
	// ノ --> の
	// ハ --> は
	// バ --> ば
	// パ --> ぱ
	// ヒ --> ひ
	// ビ --> び
	// ピ --> ぴ
	// フ --> ふ
	// ブ --> ぶ
	// プ --> ぷ
	// ヘ --> へ
	// ベ --> べ
	// ペ --> ぺ
	// ホ --> ほ
	// ボ --> ぼ
	// ポ --> ぽ
	// マ --> ま
	// ミ --> み
	// ム --> む
	// メ --> め
	// モ --> も
	// ャ --> ゃ
	// ヤ --> や
	// ュ --> ゅ
	// ユ --> ゆ
	// ョ --> ょ
	// ヨ --> よ
	// ラ --> ら
	// リ --> り
	// ル --> る
	// レ --> れ
	// ロ --> ろ
	// ヮ --> ゎ
	// ワ --> わ
	// ヰ --> ゐ
	// ヱ --> ゑ
	// ヲ --> を
	// ン --> ん
	// ヴ --> ゔ
	// ヵ --> ゕ
	// ヶ --> ゖ
}

func ExampleKanas() {
	kanas := kana.Kanas("ABCabcあぁいぃうゔぅえぇおぉやゃゆゅよょわゎかゕけゖ")

	fmt.Println("Stringer:", kanas)
	fmt.Println("String:", kanas.String())
	fmt.Println("ToHiragana:", kanas.ToHiragana())
	fmt.Println("ToKatakana:", kanas.ToKatakana())
	// Output:
	// Stringer: ABCabcあぁいぃうゔぅえぇおぉやゃゆゅよょわゎかゕけゖ
	// String: ABCabcあぁいぃうゔぅえぇおぉやゃゆゅよょわゎかゕけゖ
	// ToHiragana: ABCabcあぁいぃうゔぅえぇおぉやゃゆゅよょわゎかゕけゖ
	// ToKatakana: ABCabcアァイィウヴゥエェオォヤャユュヨョワヮカヵケヶ
}

func ExampleKanas_MarshalJSON() {
	kanas := kana.Kanas("あいうえお")

	kanjiJSON, err := kanas.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MarshalJSON:", string(kanjiJSON))
	// Output: MarshalJSON: "あいうえお"
}

func ExampleKanas_UnmarshalJSON() {
	// Note the type of the "Data" field is "Kanas"
	type dummy struct {
		Data kana.Kanas `json:"data"`
	}

	// Data in JSON format
	sampleData := []byte(heredoc.Doc(`{
		"data": "あいうえお"
	}`))

	// Parse JSON to Go object
	parsedJSON := dummy{}

	err := json.Unmarshal(sampleData, &parsedJSON)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsedJSON.Data)
	// Output: あいうえお
}
