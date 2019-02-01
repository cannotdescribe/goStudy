package board

import (
	"fmt"
)

type Board struct {
	arr []int
	//变化位
	flag int 
	stepContainer *StepContainer
}

func NewBoard(arr []int, flag int, stepContainer *StepContainer) *Board{
	return &Board{
		arr: arr,
		flag: flag,
		stepContainer: stepContainer,
	}
}

func (bo Board) CheckCell(checkName int) bool{
	x := bo.flag % 9
	for i:=0;i<9;i++{
		if(bo.arr[i*9+x] == checkName){
			return true
		}
	}
	return false
}

func (bo Board) CheckRow(checkName int) bool{
	y := bo.flag / 9
	az := bo.arr[y*9 : (y+1)*9]
	for i:=0;i<len(az);i++{
		if(az[i] == checkName){
			return true
		}
	}
	return false
}

func (bo Board) CheckPalace(checkName int) bool{
	x := bo.flag % 9
	y := bo.flag / 9
	initX := x/3 * 3
	initY := y/3 * 3
	for i:=0;i<3;i++{
		if(bo.arr[x+initY*3 +initX +i] == checkName){
			return true;
		}
	}
	for i:=0;i<3;i++{
		if(bo.arr[x+(initY+1)*3 +initX +i] == checkName){
			return true;
		}
	}
	for i:=0;i<3;i++{
		if(bo.arr[x+(initY+2)*3 +initX +i] == checkName){
			return true;
		}
	}
	return false;
}

func (bo Board) NotCheck(checkName int) bool{
	return !bo.CheckCell(checkName) && !bo.CheckRow(checkName) && !bo.CheckPalace(checkName)
}

func (bo Board) CopyBoard() *Board{
	copyArray := make([]int, 81)
	copy(copyArray, bo.arr)
	return &Board{
		arr: copyArray,
		flag: bo.flag,
	}
}

func (bo Board) Guess() (bool, *Board){
	fmt.Println(bo)
	for i:=bo.arr[bo.flag]+1; i<=9; i++ {
		if(bo.NotCheck(i)){
			copyBoard := bo.CopyBoard()
			copyBoard.arr[bo.flag] = i
			flag := bo.flag+1
			for;bo.stepContainer.isInitData(flag);{
				flag++
			}
			copyBoard.flag = flag
			return true, copyBoard
		}
	}
	return false, &bo
}

func (bo Board) IsLast() bool{
	for i:=bo.arr[bo.flag]+1; i<=9; i++ {
		if(bo.NotCheck(i)){
			return true
		}
	}
	return false
}

func (bo Board) getIndex() int{
	return bo.flag
}