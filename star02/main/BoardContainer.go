package main

import (
	"go_code/star02/utils"
)

type BoardConainer struct {
	arr [81]int
}

func (bc BoardConainer) GetIndex(index int)(x,y int){
	return utils.GetIndex(bc.arr, index)
}

func (bc BoardConainer) GetPoint(x, y int) int{
	return utils.GetPoint(bc.arr, x, y)
}

func (bc BoardConainer) CheckCell(index int, checkName int) bool{
	return utils.CheckCell(bc.arr, index, checkName)
}

func (bc BoardConainer) CheckRow(index int, checkName int) bool{
	return utils.CheckRow(bc.arr, index, checkName)
}

func (bc BoardConainer) UpGrid(index int, length int) (int, int, error) {
	return utils.UpGrid(bc.arr, index, length)
}


