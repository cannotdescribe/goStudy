package main

import (
	"fmt"
)

func sqrt12(x float64) float64{
	var z float64 = x/2;
	for i:=0; i<100;i++{
		z -= (z*z -x)/(2*z);
		if(z*z == x){
			fmt.Println("开跟结束");
			break;
		}
	}
	return z;
}

//func main(){
//	fmt.Println(sqrt12(81));
//}
