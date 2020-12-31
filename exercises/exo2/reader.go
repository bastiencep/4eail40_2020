package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (tr spaceEraser) Read(buf []byte) (int, error) {
	n, err := tr.r.Read(buf)
	j := 0
	for i := 0; i < n; i++ {
		if buf[i] != 32 {
			buf[j] = buf[i]
			j++
		}
	}
	return j, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
