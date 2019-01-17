package main

import (
	"fmt"
	"runtime"
)

func os(){
	fmt.Println(runtime.GOOS);
	switch os := runtime.GOOS; os{
		case "windows":
			fmt.Println("window.os");
		case "linux":
		    fmt.Println("linux.os");
		default:
		    fmt.Println("sometime others");
	}
}

func ot(){
	switch c :=12; {
		case c<13:
			fmt.Println("《12");
		case c>11:
			fmt.Println(">11");
		default:
		    fmt.Println("else");		
	}
}

func def() int{
	defer fmt.Println("推迟..执行");
	fmt.Println("这里");
	return 54;
}

//func main(){
//	fmt.Println("start");
//	for i:=0;i<10;i++{
//		defer fmt.Println(i);
//	}
//	fmt.Println("end");
//}

