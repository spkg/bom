// +build !go1.5

package bom

import "bufio"

func discardBytes(buf *bufio.Reader, n int) {
	// cannot use the buf.Discard method as it was introduced in Go 1.5
	if n <= 0 {
		return
	}
	b := make([]byte, n)
	buf.Read(b)
}
