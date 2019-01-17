package main

import (
//	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	resultMap := make(map[string]int);
	for _,v := range arr{
		_,contains := resultMap[v]
		if contains{
			resultMap[v] ++
		}else{
			resultMap[v] = 1;
		}
	}
	
	return resultMap
}

//func main(){
//	a := WordCount("hello I am wanghui, I from China, I like programme");
//	fmt.Println(a);
//}
