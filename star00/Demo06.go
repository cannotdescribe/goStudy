package main

import (
//	"fmt"
)

type Vertex struct{
	x int
	y int
}

var (
	v1 = Vertex{1,2};
	v2 = &Vertex{1,3}
	v3 = Vertex{x:12};
	v4 = Vertex{};
)

//func main(){
//	v2.x = 32;
//	fmt.Println(v1);
//	fmt.Println(v2);
//	fmt.Println(v3);
//	fmt.Println(v4);
//}

