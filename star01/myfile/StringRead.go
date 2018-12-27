package myfile

import (
	"fmt"
	"strings"
)

func StringReadGo() {
	r := strings.NewReader("hello world!")
	b := make([]byte, 8)
	for {
		v, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", v, err, b)
		fmt.Printf("b[:n] = %q\n", b[:v])
		if err != nil {
			break
		}
	}

}
