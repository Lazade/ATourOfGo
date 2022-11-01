package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13r.r.Read(p)
	var temp uint8
	for i := 0; i < len(p); i++ {
		bit := p[i]
		fmt.Println(i, bit)
		if bit >= 65 && bit <= 90 {
			temp = bit + 13
			if temp > 90 {
				temp = temp - 90 + 65
			}
			p[i] = temp
		} else if bit >= 97 && bit <= 122 {
			temp = bit + 13
			if temp > 122 {
				temp = temp - 122 + 97
			}
			p[i] = temp
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
