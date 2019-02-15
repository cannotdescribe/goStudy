package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func (stu *Student) setName(name string){
	stu.Name = name
}

func main(){
	student := Student{
		Name: "wanghui",
		Age: 28,
	}
	student.setName("bilibili");

	v, e := json.Marshal(student)

	fmt.Println(string(v),e)

	arr := []int{1,2,3,4,5}
	
	for i,num := range arr {
		fmt.Println(i, num)
	}
}