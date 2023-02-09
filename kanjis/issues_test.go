package kanjis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_issue1(t *testing.T) {
	for _, test := range []struct {
		input  rune
		expect rune
	}{
		// Variants of old kanji
		{'邊', '辺'},
		{'邉', '辺'},
		{'辨', '弁'},
		{'辯', '弁'},
		{'瓣', '弁'},
		// Non joyo kanjis but in the kanji.NonJoyoOld2NewMap map
		{'鬪', '闘'},
		{'鬭', '闘'},
	} {
		assert.Equal(t, test.expect, FixRuneAsJoyo(test.input),
			"'%q' should be converted to '%q'", string(test.input), string(test.expect))
	}
}
