package main

import (
//	"fmt"
//	"math"
)

func computed(fn func(float64, float64) float64) float64{
	return fn(3, 4);
}

//func main(){
//	var a = func(a,b float64) float64{
//		return math.Sqrt(a*a+b*b);
//	}
//	fmt.Println(computed(a));
//}
