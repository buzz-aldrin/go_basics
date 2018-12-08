package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		if isLower(b[i]) {
			b[i] = (((b[i] - byte(97)) + byte(13)) % byte(26)) + byte(97)
		} else if isUpper(b[i]) {
			b[i] = (((b[i] - byte(65)) + byte(13)) % byte(26)) + byte(65)
		}
	}

	return n, nil
}

func isLower(ch byte) bool {
	if ch >= byte(97) && ch <= byte(122) {
		return true
	}
	return false
}

func isUpper(ch byte) bool {
	if ch >= byte(65) && ch <= byte(90) {
		return true
	}
	return false
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
