package methodsandinterface

import (
	"golang.org/x/tour/reader"
)

type myReader struct{}

func (r myReader) Read(bytes []byte) (int, error) {

	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil

}

//TestReader needs a comment
func TestReader() {
	reader.Validate(myReader{})
}
