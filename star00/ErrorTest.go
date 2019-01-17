package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return "not -1"
}

func getErr() error{
	return ErrNegativeSqrt(0)
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		fmt.Println(getErr());
	}
	return 0, nil
}
//
//func main() {
//	fmt.Println(Sqrt(2))
//	fmt.Println(Sqrt(-2))
//}
