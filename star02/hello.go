package main

import (
	"go_code/star02/board"
	"go_code/star02/utils"
	// "fmt"
)

type DD struct{
	set *utils.Set
}

func newDD(flag int) *DD{
	mapa := make(map[interface{}]bool)
	ma := utils.NewSet(mapa)
	d := &DD{set: ma}
	
	return d
}

type Demo01 struct{
	ia []DD
	a int
}

func main() {
	sudo := []int{
		0, 1, 0, 0, 2, 0, 7, 0, 0,
		7, 0, 5, 0, 3, 4, 0, 0, 0,
		1, 0, 0, 6, 0, 0, 0, 0, 8,
		0, 3, 0, 0, 0, 0, 0, 8, 0,
		0, 7, 8, 0, 0, 0, 5, 6, 0,
		0, 0, 1, 0, 0, 0, 3, 0, 0,
		0, 6, 0, 0, 1, 0, 9, 0, 0,
		9, 0, 0, 0, 0, 3, 0, 0, 2,
		0, 0, 0, 2, 6, 0, 0, 3, 7}

	stepContainer := board.NewStepContainer(sudo)
	
	for i:=0;i<10;i++{
		if(stepContainer.NextStep()){
			stepContainer.Before()
		}
	}
}
