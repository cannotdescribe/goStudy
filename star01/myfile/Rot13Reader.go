package myfile

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	io io.Reader
}

func (reader Rot13Reader) Read(d []byte) (n int, err error) {
	v, err := reader.io.Read(d)
	for i := 0; i < len(d[:v]); i++ {
		switch {
		case d[i] > 'A' && d[i] <= 'M':
			d[i] += 13
		case d[i] > 'M' && d[i] <= 'Z':
			d[i] -= 13
		case d[i] > 'a' && d[i] <= 'm':
			d[i] += 13
		case d[i] > 'm' && d[i] <= 'z':
			d[i] -= 13
		}
	}
	return v, err
}

func Rot13ReaderFunc() {
	v := strings.NewReader("hello world!")
	r := Rot13Reader{v}
	io.Copy(os.Stdout, &r)
	io.Copy(os.Stdout, v) //没有值可以读取
}
