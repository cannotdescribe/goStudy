package board

import (
	"go_code/star02/utils"
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
		if(bo.arr[initY*9 +initX +i] == checkName){
			return true;
		}
	}
	for i:=0;i<3;i++{
		if(bo.arr[(initY+1)*9 +initX +i] == checkName){
			return true;
		}
	}
	for i:=0;i<3;i++{
		if(bo.arr[(initY+2)*9 +initX +i] == checkName){
			return true;
		}
	}
	return false;
}

//检车checkName是否允许被填写
func (bo Board) NotCheck(checkName int) bool{
	return !bo.CheckCell(checkName) && !bo.CheckRow(checkName) && !bo.CheckPalace(checkName)
}

func (bo Board) CopyBoard() *Board{
	copyArray := make([]int, 81)
	copy(copyArray, bo.arr)
	return &Board{
		arr: copyArray,
		flag: bo.flag,
		stepContainer: bo.stepContainer,
	}
}

func (bo Board) Guess() bool{
	for i:=bo.arr[bo.flag]+1; i<=9; i++ {
		if(bo.NotCheck(i)){
			bo.arr[bo.flag] = i
			return true
		}
	}
	return false
}

//到最后了,到最后了无法进行猜测
func (bo Board) IsLast(endWatch int) bool{
	for i:=endWatch+1; i<=9; i++ {
		if(bo.NotCheck(i)){
			return false
		}
	}
	return true
}

func (bo Board) getIndex() int{
	return bo.flag
}

func (bo Board) LastValue() int{
	return bo.arr[bo.flag]
}

func (bo Board) LastSecondValue(setFlag *utils.Set) (int, int){
	flag := bo.flag-1
	for ;setFlag.Has(flag); {
		flag--
	}
	if flag == -1 {
		return -1, -1
	}else{
		return flag, bo.arr[flag]
	}
}