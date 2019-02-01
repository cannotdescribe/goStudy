package board

import (
	"go_code/star02/utils"
	"fmt"
)

type StepContainer struct{
	bos *[]Board
	flag *utils.Set
}

func NewStepContainer(sodu []int) *StepContainer{
	bo := make([]Board, 0, 81)
	stepContainer := &StepContainer{
		bos: &bo,
		flag: utils.NewSet(map[interface{}]bool{}),
	}
	*stepContainer.bos = append(*stepContainer.bos, *NewBoard(sodu, 0, stepContainer))
	fmt.Println("stepContainer", stepContainer)
	for index, value := range sodu{
		if(value != 0){
			stepContainer.flag.Add(index)
		}
	}
	return stepContainer
}

func (sc *StepContainer) isInitData(index int) bool{
	return sc.flag.Has(index)
}

func (sc *StepContainer) appendBoard(board *Board){
	*sc.bos = append(*sc.bos, *board)
}

func (sc *StepContainer) NextStep() bool{
	index := len(*sc.bos)
	
	flag, board := (*sc.bos)[index-1].Guess()
	if(flag){
		sc.appendBoard(board)
	}
	return flag
}

func (sc *StepContainer) Before(){
	*sc.bos = (*sc.bos)[:len(*sc.bos)-1]
	flag := len(*sc.bos)-1
	for ;(*sc.bos)[flag].IsLast();{
		flag--
		*sc.bos = (*sc.bos)[:len(*sc.bos)-1]
	}
}
