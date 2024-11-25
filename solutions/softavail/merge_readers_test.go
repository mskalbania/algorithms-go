package softavail

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestMergeReadersWorksCorrectly(t *testing.T) {
	testData := []struct {
		name           string
		r1, r2         io.Reader
		expectedResult []byte
	}{
		{
			"empty readers",
			strings.NewReader(""),
			strings.NewReader(""),
			nil,
		},
		{
			"first empty, second not",
			strings.NewReader(""),
			strings.NewReader("abc"),
			nil,
		},
		{
			"second empty, first not",
			strings.NewReader("abc"),
			strings.NewReader(""),
			nil,
		},
		{
			"same length",
			strings.NewReader("abc"),
			strings.NewReader("ABC"),
			[]byte("aAbBcC"),
		},
		{
			"different length",
			strings.NewReader("abcd"),
			strings.NewReader("abc"),
			[]byte("aabbcc"),
		},
	}
	for _, datum := range testData {
		t.Run(datum.name, func(t *testing.T) {
			resultReader := MergeReaders(datum.r1, datum.r2)
			buffer := bytes.Buffer{}
			_, err := io.Copy(&buffer, resultReader)

			require.NoError(t, err)
			require.Equal(t, datum.expectedResult, buffer.Bytes())
		})
	}
}
