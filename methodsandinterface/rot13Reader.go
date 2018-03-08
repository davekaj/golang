package methodsandinterface

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	size, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i, v := range b {
		if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
			b[i] = v + 13
		} else if v >= 'n' && v <= 'z' || v >= 'N' && v <= 'Z' {
			b[i] = v - 13
		}
	}

	return size, nil
}

//Rot13Export needs a comment
func Rot13Export() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
