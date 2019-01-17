package main

import (
	"fmt"
)

func printArray(a []int){
	fmt.Printf("val: %v. len: %v, cap: %v \n", a, len(a), cap(a));
}

//func main(){
//	var a []int = []int{1,2,3,4,5,6,7,8};
//	b := a[:2];
//	c := a[1:];
//	printArray(a);
//	printArray(b);
//	
//	b[4] = 233;
//	printArray(b);
//	
//	printArray(c);
//}
