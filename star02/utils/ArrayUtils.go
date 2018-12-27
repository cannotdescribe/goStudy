package utils

import (
	"go_code/star02/myError"
)

func CheckRow(arr [81]int, index int, checkName int) bool{
	x := index / 9
	az := arr[x*9 : (x+1)*9]
	for i:=0;i<len(az);i++{
		if(az[i] == checkName){
			return true
		}
	}
	return false
}

func UpGrid(arr [81]int, index int, length int) (int, int, error) {
	upIndex := index - 9*length
	if upIndex < 0 {
		return 0, 0, myError.IndexOverFlowError{}
	} else {
		return arr[upIndex], upIndex, myError.IndexOverFlowError{}
	}
}