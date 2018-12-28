package utils

import (
	"fmt"
	"go_code/star02/myError"
)

func GetIndex(arr [81]int, index int)(x,y int){
	return index%9, index/9
}

func GetPoint(arr [81]int, x, y int) int{
	return y*9+x
}

func CheckCell(arr [81]int, index int, checkName int) bool{
	x := index % 9
	for i:=0;i<9;i++{
		if(arr[i*9+x] == checkName){
			return true
		}
	}
	return false
}

func CheckRow(arr [81]int, index int, checkName int) bool{
	y := index / 9
	az := arr[y*9 : (y+1)*9]
	for i:=0;i<len(az);i++{
		if(az[i] == checkName){
			return true
		}
	}
	return false
}

func CheckPalace(arr [81]int, index int, checkName int) bool{
	x := index % 9
	y := index / 9
	initX := x/3 * 3
	initY := y/3 * 3
	fmt.Println(x,y, initX, initY)
	return false;
}

func UpGrid(arr [81]int, index int, length int) (int, int, error) {
	upIndex := index - 9*length
	if upIndex < 0 {
		return 0, 0, myError.IndexOverFlowError{}
	} else {
		return arr[upIndex], upIndex, myError.IndexOverFlowError{}
	}
}
