package main

//import (
//	"fmt"
//)

func add(x, y int) int {
	return x+y;
}

func swap(x, y string) (string, string){
	return y,x;
}

func split(sum int) (x, y int){
	x = sum/10;
	y = sum%10;
	return;
}
//func main(){
//	fmt.Println("hello world", rand.Intn(20));
//	fmt.Println(add(1,2));
//	a ,b := swap("甲", "乙");
//	fmt.Println(a, b);
//	a, b :=split(12);
//	fmt.Println(a, b);
//}

