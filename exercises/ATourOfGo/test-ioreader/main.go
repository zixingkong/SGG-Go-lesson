
package main

import (
"io"
"os"
"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	var b byte
	for i := 0; i < n; i++ {
		if c :=p[i]; c >= 'A' && c <= 'z' {
			if c >= 'a' {
				b = byte('a')
			} else {
				b = byte('A')
			}
			// rotate by 13 places
			p[i] = (c - b + 13) % 26 + b
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
