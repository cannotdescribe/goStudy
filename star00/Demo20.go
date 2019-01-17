package main

import (
//	"fmt"
)

func adder() func(int) int {
	sum := 0;
	return func(a int) int{
		sum += a;
		return sum;
	}
}

//func main(){
//	a, b :=adder(), adder()
//	
//	for i:=0;i<100;i++ {
//		fmt.Println(
//			a(i),
//			b(i),
//		)
//	}
//}