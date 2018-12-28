package main

import (
	"fmt"
	"go_code/star02/utils"
)

func main() {
	val := [81]int{
		0, 1, 0, 0, 2, 0, 7, 0, 0,
		7, 0, 5, 0, 3, 4, 0, 0, 0,
		1, 0, 0, 6, 0, 0, 0, 0, 8,
		0, 3, 0, 0, 0, 0, 0, 8, 0,
		0, 7, 8, 0, 0, 0, 5, 6, 0,
		0, 0, 1, 0, 0, 0, 3, 0, 0,
		0, 6, 0, 0, 1, 0, 9, 0, 0,
		9, 0, 0, 0, 0, 3, 0, 0, 2,
		0, 0, 0, 2, 6, 0, 0, 3, 7}

	result := utils.CheckRow(val, 20, 6)
	fmt.Println(result)

	num := utils.CheckPalace(val, 20, 2);
	fmt.Println(num)
}
