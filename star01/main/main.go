// start-demo01 project main.go
package main

import (
	"gobase"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	gobase.ByteReadD()

}
