package gobase

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

func mcWriteAt() {
	file, _ := os.Create("text.txt")
	defer file.Close()
	dir, _ := filepath.Abs("test.txt")
	fmt.Println(dir)
	file.Write([]byte("我测试时的这段话"))

	file.WriteAt([]byte("尔尔"), 6)
}

func mcReadAt() {
	file, _ := os.Open("text.txt")
	defer file.Close()
	b := make([]byte, 30)
	n, err := file.ReadAt(b, 6)
	fmt.Printf("%q \n", b[:n])
	fmt.Println(n, err)
}

func readData() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	write := bufio.NewWriter(os.Stdout)
	write.ReadFrom(file)
	write.Flush()
}

func readData_d1() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	read := bufio.NewReader(file)
	b := make([]byte, 30)

	n, _ := read.Read(b)
	fmt.Printf("%q \n", b[:n])

}

func readData_d2() {
	file, _ := os.Open("text.txt")
	defer file.Close()
	write := bufio.NewWriter(os.Stdout)
	write.ReadFrom(file)
	b := make([]byte, 30)

	n, _ := write.Write(b)
	fmt.Printf("%q \n", b[:n])
	fmt.Println(n)
	write.Flush()
}

func writeData() {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

func seekD() {
	reader := strings.NewReader("acfun是个视频网站")
	reader.Seek(-9, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c \n", r)
	z, _, _ := reader.ReadRune()
	fmt.Printf("%c \n", z)
}

/**
 * 无法进测试，只能通过main方法运行才能正常工作
 */
func ByteReadD() {
	//	var buff bytes.Buffer = bytes.Buffer()
	var buffer *bytes.Buffer = new(bytes.Buffer)
	var ch byte
	fmt.Scanf("%c\n", &ch)

	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入成功")
		b, _ := buffer.ReadByte()
		fmt.Printf("写入的值为%c \n", b)
	} else {
		fmt.Println("写入失败")
	}

	fmt.Println(buffer)
}

func ByteReadD2() {
	var bu *bytes.Buffer = new(bytes.Buffer)
	bu.Write([]byte{'a', 'b', 'c'})
	b1, _ := bu.ReadByte()
	fmt.Printf("%c \n", b1)
	b2, _ := bu.ReadByte()
	fmt.Printf("%c \n", b2)

	err := bu.UnreadByte()
	b3, _ := bu.ReadByte()
	fmt.Printf("%v %c \n", err, b3)
}

func newByteBufferD1() {
	var bu *bytes.Buffer = bytes.NewBuffer([]byte{'a', 'b', 'c'})
	b1, _ := bu.ReadByte()
	fmt.Printf("%q \n", b1)
	b2, _ := bu.ReadByte()
	fmt.Printf("%q \n", b2)
	bu.UnreadByte()
	b3, _ := bu.ReadByte()
	fmt.Printf("%q \n", b3)
}

func stringIndexTest() {
	st := "测试一1二a三"
	index := strings.Index(st, "三")
	fmt.Println(index)
	a := utf8.RuneCountInString(st[:index])
	fmt.Println(a)
}

func sectionReadD1() {
	read := strings.NewReader("测试一1二a三")
	var sr *io.SectionReader = io.NewSectionReader(read, 6, 7)
	b := make([]byte, 20)
	n, _ := sr.Read(b)
	fmt.Printf("%q \n", b[:n])
}

func limitReadD1() {
	read := strings.NewReader("测试一1二a三")
	//	var lr io.LimitedReader = io.LimitedReader{read, 10}

	var lr io.Reader = io.LimitReader(read, 10)
	b := make([]byte, 20)
	n, _ := lr.Read(b)
	fmt.Printf("%q \n", b[:n])
}

func pipeWriteD1(write *io.PipeWriter) {
	for i := 0; i < 3; i++ {
		_, err := write.Write([]byte(fmt.Sprintf("%v%v", "对的", i)))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func pipeReadD1(read *io.PipeReader) {
	for i := 0; i < 3; i++ {
		b := make([]byte, 20)
		n, err := read.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%q \n", b[:n])
	}
	fmt.Println("end!")
}

func pipeMainD1() {
	read, write := io.Pipe()
	go pipeReadD1(read)
	go pipeWriteD1(write)
	time.Sleep(10 * time.Second)
}

func readformCopy() {
	read := bufio.NewReader(strings.NewReader("测试哈哈哈哈"))
	n, err := io.Copy(os.Stdout, read)
	fmt.Println(n, err)
}

func readformCopyN() {
	read := bufio.NewReader(strings.NewReader("测试哈哈哈哈"))
	n, err := io.CopyN(os.Stdout, read, 21)
	fmt.Println(n, err)
}

func arrayTest(a []int) {
	a[0] = 1
}
func arrayTD() {
	//	a := make([]int, 5)
	a := []int{0, 0}
	arrayTest(a)
	fmt.Println(a)
}

func baseTypeD_r(a *int) {
	*a = 10
}

func baseTypeD() {
	b := 1
	baseTypeD_r(&b)
	fmt.Println(b)
}

func readAtLeastD() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	b := make([]byte, 40)
	n, err := io.ReadAtLeast(file, b, 25)
	fmt.Println(err)
	fmt.Printf("%q \n", b[:n])
}

func readFullD() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, 40)
	n, err := io.ReadFull(file, b)
	fmt.Println(err)
	fmt.Printf("%q \n", b[:n])
}

func writeStringD() {
	file, err := os.OpenFile("text.txt", os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	io.WriteString(file, "input")
}

func multiD() {
	f1, _ := os.Open("text.txt")
	f2, _ := os.Open("text_f.txt")
	multiReader := io.MultiReader(f1, f2)
	b := make([]byte, 6)
	result := make([]byte, 0, 120)
	for f, err := multiReader.Read(b); err != io.EOF; f, err = multiReader.Read(b) {
		if err != nil {
			return
		}
		fmt.Printf("%q \n", b[:f])
		result = append(result, b[:f]...)
	}
	fmt.Printf("%q", result)
}

func NopCloser NopCloseD(){
	
	
}
