package main

import (
//	"fmt"
)

type Pep struct{
	X, Y float64;
}

func (p *Pep)Scale(f float64){
	p.X = f * p.X;
	p.Y = f * p.Y;
}

func (p Pep)Scale11(f float64){
	p.X = f * p.X;
	p.Y = f * p.Y;
}

func Scale_(p *Pep, f float64){
	p.X = f * p.X;
	p.Y = f * p.Y;
}

func Scale_11(p Pep, f float64){
	p.X = f * p.X;
	p.Y = f * p.Y;
}

//func main(){
//	p := Pep{1, 2};
//	p.Scale(10);
//	fmt.Println(p);
//	Scale_(&p, 10);
//	fmt.Println(p);
//	
//	(&p).Scale11(10);
//	fmt.Println(p);
//	Scale_11(p, 10);
//	fmt.Println(p);
//}
