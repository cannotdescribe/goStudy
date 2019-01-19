package board

import (
//	"go_code/star02/utils"
)

type Board struct {
	arr [81]int
	//变化位
	flag int 
}

func NewBoard(arr [81]int, flag int) *Board{
	return &Board{
		arr: arr,
		flag: flag,
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

func (bo Board) Check(checkName int) bool{
	return bo.CheckCell(checkName) &&  bo.CheckRow(checkName) &&  bo.CheckPalace(checkName)
}

func (bo Board) getIndex() int{
	return bo.flag
}