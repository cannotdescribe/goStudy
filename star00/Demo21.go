package main

import (
//	"fmt"
	"math"
)
type AC struct{
	X, Y float64
} 

func (p AC)Abs() float64{
	return math.Sqrt(p.X * p.X + p.Y * p.Y);
}

type MyFloat float64;

func (f MyFloat)Abs() float64{
	if f < 0{
		return -float64(f);
	}
	return float64(f);
} 

//func main(){
//	ac := AC{3,4};
//	fmt.Println(ac.Abs());
//	f := MyFloat(-math.Sqrt2);
//	fmt.Println(f.Abs());
//}