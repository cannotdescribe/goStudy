package thread

import (
	"fmt"
	"time"
)

func world() {
	fmt.Println("go")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("world!")
	}
}

func GoThread() {
	fmt.Print("hello, ")
	go world()
	time.Sleep(1000 * time.Millisecond)
}

func sum(g []int, c chan int) {
	s := 0
	for _, v := range g {
		s += v
	}
	c <- s
}

func GoChannel() {
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	go sum(g[len(g)/2:], c)
	go sum(g[:len(g)/2], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}

func TtChan() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Print("over heee")
	fmt.Println(<-c)
}
