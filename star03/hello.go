package main

import "fmt"

type Peo struct{
	oo *[]int
}

func ac(po Peo){
	*po.oo = append(*po.oo, 10)
}

func main(){
	var aq []int = make([]int, 0)
	a := Peo{
		oo: &aq,
	}
	ac(a)
	fmt.Println(*a.oo)
}