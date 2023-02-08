package kana

// Kanas is a type that represents a string of Kana characters.
type Kanas []rune

// String is a Stringer interface implementation.
func (k Kanas) String() string {
	return string(k)
}

// MarshalJSON is a Marshaler interface implementation.
func (k Kanas) MarshalJSON() ([]byte, error) {
	return []byte("\"" + string(k) + "\""), nil
}

// UnmarshalJSON is a Unmarshaler interface implementation.
func (k *Kanas) UnmarshalJSON(data []byte) error {
	// Remove the double quotes in JSON
	data = data[1 : len(data)-1]

	*k = Kanas(string(data))

	return nil
}

// ToHiragana converts the data to Hiragana.
func (k Kanas) ToHiragana() *Kanas {
	for i, r := range k {
		k[i] = ToHiragana(r)
	}

	return &k
}

// ToKatakana converts the data to Katakana.
func (k Kanas) ToKatakana() *Kanas {
	for i, r := range k {
		k[i] = ToKatakana(r)
	}

	return &k
}
