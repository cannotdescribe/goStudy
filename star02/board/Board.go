package board

import (
//	"go_code/star02/utils"
)

type Board struct {
	arr [81]int
	//变化位
	flag int 
}

func (bo Board) CheckCell(index int, checkName int) bool{
	x := index % 9
	for i:=0;i<9;i++{
		if(bo.arr[i*9+x] == checkName){
			return true
		}
	}
	return false
}

func (bo Board) CheckRow(index int, checkName int) bool{
	y := index / 9
	az := bo.arr[y*9 : (y+1)*9]
	for i:=0;i<len(az);i++{
		if(az[i] == checkName){
			return true
		}
	}
	return false
}

func (bo Board) CheckPalace(index int, checkName int) bool{
	x := index % 9
	y := index / 9
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

func (bo Board) Check(index int, checkName int) bool{
	return bo.CheckCell(index, checkName) &&  bo.CheckRow(index, checkName) &&  bo.CheckPalace(index, checkName)
}
