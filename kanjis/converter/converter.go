package converter

import (
	"io"

	"github.com/pkg/errors"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

// ----------------------------------------------------------------------------
//  Type: Converter
// ----------------------------------------------------------------------------

// Converter is a wrapper of transform.Transformer. It provides an additional
// method `Convert` to convert the input to the output.
type Converter struct {
	runeTransformer runes.Transformer
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a new Converter object. The given function will be used to convert
// the characters.
// Note that if the given function returns -1, the character will be omitted.
func New(fn func(in rune) rune) Converter {
	return Converter{
		runeTransformer: runes.Map(fn),
	}
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

func (trns *Converter) Convert(input io.Reader, output io.Writer) error {
	if input == nil || output == nil {
		return errors.New("input or output is nil")
	}

	_, err := io.Copy(output, transform.NewReader(input, trns.runeTransformer))

	return errors.Wrap(err, "failed to copy the input to the output")
}
