package main

import (
	"fmt"
)

//func main(){
//	var z interface{} = "hello";
//	printVal(z);
//	a, _ := z.(string);
//	printType(a);
//	
//	printType(true);
	
//	printVal(a, ak);
//	b, bk := z.(float64);
//	printVal(b, bk);
//}

func printType(z ...interface{}){
	for _,v := range z{
		switch vv := v.(type){
			case int:
			    fmt.Printf("int %v \n", vv);
			case string:
			    fmt.Printf("string %v \n", vv);
			default:
			    fmt.Printf("type: %T", vv);
		}
	}
}

func printVal(z ...interface{}){
	for _,v := range z{
		fmt.Printf("val: %v, type: %T \n", v, v);
	}
} 


