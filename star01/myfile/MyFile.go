package myfile

import (
	"fmt"
	"io"
	"os"
)

func Read(src string, dest string) (written int64, err error) {
	fileDest, err := os.OpenFile(dest, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fileDest, err = os.Create(dest)
		fmt.Println(fileDest, err, err != nil)

		if err != nil {
			return
		}
	}
	defer fileDest.Close()

	fileSrc, e := os.OpenFile(src, os.O_RDWR|os.O_APPEND, 0660)
	//	fileSrc, e := os.Open(src)
	fmt.Println(fileSrc, e)
	fmt.Println(fileDest)
	if e != nil {
		return
	}
	defer fileSrc.Close()
	fmt.Println("copy 完成")
	return io.Copy(fileDest, fileSrc)
}

func Hello() {
	fmt.Println("fmt")
}
