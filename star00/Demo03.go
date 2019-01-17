package main

import (
	"fmt"
	"math"
)

func sqrt(i int) string{
	if i<0{
		return sqrt(-i) +"i";
	}else{
		return fmt.Sprint(math.Sqrt(float64(i)));
	}
}

func pow(a, b, c float64) float64{
	if v := math.Pow(a, b); v<c{
		return v; 
	}
	return c;
}

//func main(){
//	i := -12;
//	fmt.Println(sqrt(i));
//	var i int = 12;
//	fmt.Println(-i);
//	fmt.Println(pow(3,2,10), pow(3,3,20));
//}
