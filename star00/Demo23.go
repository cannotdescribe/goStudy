package main

import (
	"fmt"
)

type People interface{
	setName(name string)
	getName() string
}

type Student struct{
	name string;
}

func (p *Student) setName(name string) {
	if p == nil{
		fmt.Println("<nil/>");
		return;
	}
	p.name = name;
}

func (p *Student) getName() string {
	return p.name;
}

//func main(){
//	var he People;
//	var stu *Student;
//	he = stu;
//	he.setName("dasdas");
//	fmt.Println(he);
//	fmt.Println(he.getName());
//}

