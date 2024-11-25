package softavail

import (
	"bytes"
	"io"
)

// MergeReaders Merges readers so that they read on byte at the time.
// If any reader contains more data trim to the lower one.
// Example: r1 - abc, r2 - ABCD, result - aAbBcC
func MergeReaders(r1, r2 io.Reader) io.Reader {
	b1 := make([]byte, 1)
	b2 := make([]byte, 1)
	buffer := bytes.Buffer{}
	for {
		_, err := r1.Read(b1)
		if err == io.EOF {
			break
		}
		_, err = r2.Read(b2)
		if err == io.EOF {
			break
		}
		buffer.Write(b1)
		buffer.Write(b2)
	}
	return &buffer
}
