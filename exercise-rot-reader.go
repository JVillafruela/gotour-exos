package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	//fmt.Printf("\nn = %v err = %v  len(b)=%v\n", n, err, len(b))
	for i := 0; i < n && i < len(b); i++ {
		if 'A' <= b[i] && b[i] <= 'Z' {
			if b[i] >= 'A'+13 {
				b[i] = b[i] - 13
			} else {
				b[i] = b[i] + 13
			}
		}

		if 'a' <= b[i] && b[i] <= 'z' {
			if b[i] >= 'a'+13 {
				b[i] = b[i] - 13
			} else {
				b[i] = b[i] + 13
			}
		}

	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
