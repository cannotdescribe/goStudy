package myfile

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadAtTc_A() {
	a := strings.NewReader("he测试这个")
	b := make([]byte, 6)
	a.ReadAt(b, 5)
	fmt.Printf("%q", b)
}

func CopyFile() {
	a, _ := os.OpenFile("F:/WorkspaceGo/start-demo01/resources/ceshi.png", os.O_RDWR|os.O_APPEND, 0660)
	defer a.Close()

	b := make([]byte, 1024)

	c, _ := os.Create("F:/WorkspaceGo/start-demo01/resources/aa.png")

	defer c.Close()
	for {
		l, err := a.Read(b)
		c.Write(b[:l])

		if err == io.EOF {
			break
		}
	}

}
