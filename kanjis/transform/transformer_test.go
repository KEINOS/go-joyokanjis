package transform

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransformer(t *testing.T) {
	rot13 := func(r io.Reader) *Transformer {
		buf := make([]byte, rand.Int()%256+1) // used to test varying slice sizes

		return NewTransformer(func() ([]byte, error) {
			n, err := r.Read(buf)
			if err != nil {
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

	// simple
	msg := "Hello\n13th Floor"
	data, err := io.ReadAll(rot13(rot13(bytes.NewBufferString(msg))))

	require.NoError(t, err)
	require.Equal(t, msg, string(data))

	// random
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, 10000)

	for i := 0; i < 1000; i++ {
		_, err := rand.Read(buf)
		require.NoError(t, err, "failed to generate random data")

		data, err := io.ReadAll(rot13(rot13(bytes.NewBuffer(buf))))

		require.NoError(t, err, "failed to read Rot13-ed data from the buffer")
		require.Equal(t, string(buf), string(data))
	}
}
