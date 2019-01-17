package main

import (
	"fmt"
	"strings"
	"io"
)

func main(){
	r := strings.NewReader("hello, world!");
	b := make([]byte, 8);
	for {
		i, err := r.Read(b);
		fmt.Printf("n = %v err = %v b = %v\n", i, err, b)
		fmt.Printf("b[:n] = %q\n", b[:i])
		if(err == io.EOF){
			break;
		}
	}
}

