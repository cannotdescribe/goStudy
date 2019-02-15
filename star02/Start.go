package main

import (
	"go_code/star02/board"
	"fmt"
)

func main() {
	// sudo := []int{
	// 	0, 1, 0, 0, 2, 0, 7, 0, 0,
	// 	7, 0, 5, 0, 3, 4, 0, 0, 0,
	// 	1, 0, 0, 6, 0, 0, 0, 0, 8,
	// 	0, 3, 0, 0, 0, 0, 0, 8, 0,
	// 	0, 7, 8, 0, 0, 0, 5, 6, 0,
	// 	0, 0, 1, 0, 0, 0, 3, 0, 0,
	// 	0, 6, 0, 0, 1, 0, 9, 0, 0,
	// 	9, 0, 0, 0, 0, 3, 0, 0, 2,
	// 	0, 0, 0, 2, 6, 0, 0, 3, 7}

	// sudo := []int{
	// 	0, 3, 0, 0, 0, 5, 2, 7, 0,
	// 	0, 0, 0, 0, 0, 9, 0, 4, 0,
	// 	0, 7, 0, 4, 8, 0, 3, 0, 5,
	// 	2, 0, 0, 0, 0, 8, 0, 1, 7,
	// 	0, 0, 0, 9, 4, 6, 0, 0, 0,
	// 	9, 8, 0, 2, 0, 0, 0, 0, 6,
	// 	7, 0, 1, 0, 9, 4, 0, 3, 0,
	// 	0, 9, 0, 7, 0, 0, 0, 0, 0,
	// 	0, 5, 4, 6, 0, 0, 0, 9, 0}

	sudo := []int{
		0, 0, 0, 0, 8, 0, 3, 0, 0,
		0, 0, 9, 0, 0, 6, 0, 0, 0,
		1, 4, 0, 0, 0, 0, 0, 0, 0,
		2, 6, 8, 0, 0, 1, 9, 0, 0,
		0, 0, 0, 2, 6, 5, 0, 0, 0,
		0, 0, 7, 9, 0, 0, 1, 2, 6,
		0, 0, 0, 0, 0, 0, 0, 5, 4,
		0, 0, 0, 3, 0, 0, 7, 0, 0,
		0, 0, 2, 0, 5, 0, 0, 0, 0}

	stepContainer := board.NewStepContainer(sudo)
	i := 0
	for ;true;{
		i++;
		if(!stepContainer.NextStep()){
			if(!stepContainer.Before()){
				fmt.Println("数独不规范，或者程序有bug。")
				break;
			}
		}
		if(stepContainer.End()){
			fmt.Printf("猜测结束打印结果, 共猜测: %v 次\n", i)
			fmt.Println("结果如下: ", stepContainer.EndBoard())
			break;
		}
	}
}
