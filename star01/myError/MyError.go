package myError

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("timer: %v, what: %v.", e.When, e.What)
}

func getError() error {
	return &MyError{
		time.Now(),
		"我并不在乎这些东西",
	}
}

func ToString() {
	fmt.Println("输出数据")
}
