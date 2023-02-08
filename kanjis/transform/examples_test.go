package transform_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/KEINOS/go-joyokanjis/kanjis/transform"
	"github.com/MakeNowJust/heredoc"
)

func ExampleTransformer_lineMatcherRegExp() {
	// Filter lines matching a pattern
	matcher := func(r io.Reader, pattern string) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			for {
				line, err := br.ReadBytes('\n')
				if err != nil {
					return nil, err
				}

				matched, _ := regexp.Match(pattern, line)
				if matched {
					return line, err
				}
			}
		})
	}

	logs := heredoc.Doc(`
		23 Apr 17:32:23.604 [INFO] DB loaded in 0.551 seconds
		23 Apr 17:32:23.605 [WARN] Disk space is low
		23 Apr 17:32:23.054 [INFO] Server started on port 7812
		23 Apr 17:32:23.141 [INFO] Ready for connections
	`)

	// Pass the string though the transformer.
	out, err := io.ReadAll(matcher(bytes.NewBufferString(logs), "WARN"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// Output:
	// 23 Apr 17:32:23.605 [WARN] Disk space is low
}

// This example shows how to pipe multiple transformers together.
func ExampleTransformer_pipeline() {
	// Filter lines matching a pattern
	matcher := func(r io.Reader, pattern string) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			for {
				line, err := br.ReadBytes('\n')
				if err != nil {
					return nil, err
				}

				matched, _ := regexp.Match(pattern, line)
				if matched {
					return line, err
				}
			}
		})
	}

	// Trim space from all lines
	trimmer := func(r io.Reader) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			for {
				line, err := br.ReadBytes('\n')
				if err != nil {
					return nil, err
				}

				if len(line) > 0 {
					return append(bytes.TrimSpace(line), '\n'), nil
				}
			}
		})
	}

	// Convert a string to uppper case. Unicode aware. In this example
	// we only process one rune at a time. It works but it's not ideal
	// for production.
	toUpper := func(r io.Reader) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			c, _, err := br.ReadRune()
			if err != nil {
				return nil, err
			}

			return []byte(strings.ToUpper(string([]rune{c}))), nil
		})
	}

	phrases := "  lacy timber \n"
	phrases += "\t\thybrid gossiping\t\n"
	phrases += " coy radioactivity\n"
	phrases += "rocky arrow  \n"

	// create a transformer that matches lines on the letter 'o', trims the
	// space from the lines, and transforms to upper case.
	r := toUpper(
		trimmer(
			matcher(bytes.NewBufferString(phrases), "o"),
		),
	)

	// Pass the string though the transformer.
	out, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// Output:
	// HYBRID GOSSIPING
	// COY RADIOACTIVITY
	// ROCKY ARROW
}

func ExampleTransformer_rot13() {
	// Rot13 transformation
	rot13 := func(r io.Reader) *transform.Transformer {
		buf := make([]byte, 256)

		return transform.NewTransformer(func() ([]byte, error) {
			n, err := r.Read(buf)
			if err != nil {
				// the error including EOF will be handled by the caller
				// io.ReadAll(). So just return it.
				return nil, err
			}

			for i := 0; i < n; i++ {
				if buf[i] >= 'a' && buf[i] <= 'z' {
					buf[i] = ((buf[i] - 'a' + 13) % 26) + 'a'
				} else if buf[i] >= 'A' && buf[i] <= 'Z' {
					buf[i] = ((buf[i] - 'A' + 13) % 26) + 'A'
				}
			}

			return buf[:n], nil
		})
	}

	// Pass the string though the transformer.
	out, err := io.ReadAll(rot13(bytes.NewBufferString("Hello World")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// Output:
	// Uryyb Jbeyq
}

func ExampleTransformer_toUpper() {
	// Convert a string to uppper case. Unicode aware.
	toUpper := func(r io.Reader) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			c, _, err := br.ReadRune()
			if err != nil {
				// the error including EOF will be handled by the caller
				// io.ReadAll(). So just return it.
				return nil, err
			}

			return []byte(strings.ToUpper(string([]rune{c}))), nil
		})
	}

	// Pass the string though the transformer.
	out, err := io.ReadAll(toUpper(bytes.NewBufferString("Hello World")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// Output:
	// HELLO WORLD
}

func ExampleTransformer_trimmer() {
	// Trim space from all lines
	trimmer := func(r io.Reader) *transform.Transformer {
		br := bufio.NewReader(r)

		return transform.NewTransformer(func() ([]byte, error) {
			for {
				line, err := br.ReadBytes('\n')
				if err != nil {
					// the error including EOF will be handled by the caller
					// io.ReadAll(). So just return it.
					return nil, err
				}

				if len(line) > 0 {
					return append(bytes.TrimSpace(line), '\n'), nil
				}
			}
		})
	}

	phrases := "  lacy timber \n"
	phrases += "\t\thybrid gossiping\t\n"
	phrases += " coy radioactivity\n"
	phrases += "rocky arrow  \n"

	// Pass the string though the transformer.
	out, err := io.ReadAll(trimmer(bytes.NewBufferString(phrases)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// Output:
	// lacy timber
	// hybrid gossiping
	// coy radioactivity
	// rocky arrow
}
