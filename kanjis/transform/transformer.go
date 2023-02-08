//	Package transform provides a convenient utility for transforming one data
//	format to another.
//
// ----------------------------------------------------------------------------
// It is a part of the github.com/KEINOS/go-joyokanjis module.
//
//	https://github.com/KEINOS/go-joyokanjis/kanjis
//
// Which is originally taken from:
//
//	https://github.com/tidwall/transform
//	(please give a star to the original author if you liked this package)
//
// Due to the maintenance reasons, the file is copied here and modified to fit
// the package.
// ----------------------------------------------------------------------------
// ISC License
//
// # Copyright (c) 2017, Joshua J Baker
//
// Permission to use, copy, modify, and/or distribute this software for any purpose
// with or without fee is hereby granted, provided that the above copyright notice
// and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
// OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.
// ----------------------------------------------------------------------------
package transform

// ----------------------------------------------------------------------------
//  Type: Transformer
// ----------------------------------------------------------------------------

// Transformer represents a transform reader.
// It provides a convenient utility for transforming one data format to another.
//
// Note: This is a modified version of the original "transform" package:
// https://github.com/tidwall/transform (Copyright (c) 2017, Joshua J Baker)
type Transformer struct {
	tfn func() ([]byte, error) // user-defined transform function
	buf []byte                 // read buffer
	idx int                    // read buffer index
	err error                  // last error
}

// NewTransformer returns an object that can be used for transforming one data
// formant to another. The param is a function that performs the conversion and
// returns the transformed data in chunks/messages.
func NewTransformer(fn func() ([]byte, error)) *Transformer {
	return &Transformer{tfn: fn}
}

// ReadMessage allows for reading a one transformed message at a time.
func (r *Transformer) ReadMessage() ([]byte, error) {
	return r.tfn()
}

// Read is an implementation to conform the io.Reader interface.
func (r *Transformer) Read(p []byte) (n int, err error) {
	if len(r.buf)-r.idx > 0 {
		// There's data in the read buffer, return it prior to returning errors
		// or reading more messages.
		if len(r.buf)-r.idx > len(p) {
			// The input slice is smaller than the read buffer, copy a subslice
			// of the read buffer and increase the read index.
			copy(p, r.buf[r.idx:r.idx+len(p)])
			r.idx += len(p)

			return len(p), nil
		}

		// Copy the entire read buffer to the input slice.
		n = len(r.buf) - r.idx
		copy(p[:n], r.buf[r.idx:])

		r.buf = r.buf[:0] // reset the read buffer, keeping it's capacity
		r.idx = 0         // rewind the read buffer index

		return n, nil
	}

	if r.err != nil {
		return 0, r.err
	}

	var msg []byte

	msg, r.err = r.ReadMessage()

	// We should immediately append the incoming message to the read
	// buffer to allow for the implemented transformer to repurpose
	// it's own message space if needed.
	r.buf = append(r.buf, msg...)

	return r.Read(p)
}
