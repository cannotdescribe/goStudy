package board

import (
	"go_code/star02/utils"
	"fmt"
)

type StepContainer struct{
	bos []Board
	flag *utils.Set
}

func NewStepContainer(sodu []int) *StepContainer{
	stepContainer := &StepContainer{
		bos: make([]Board, 0, 81),
		flag: utils.NewSet(map[interface{}]bool{}),
	}
	stepContainer.bos = append(stepContainer.bos, *NewBoard(sodu, 0, stepContainer))
	fmt.Println("stepContainer", stepContainer)
	for index, value := range sodu{
		if(value != 0){
			stepContainer.flag.Add(index)
		}
	}
	return stepContainer
}

func (sc StepContainer) isInitData(index int) bool{
	return sc.flag.Has(index)
}

func (sc StepContainer) appendBoard(board *Board){
	sc.bos = append(sc.bos, *board)
}

func (sc StepContainer) NextStep() bool{
	index := len(sc.bos)
	fmt.Println("sc.bos[index-1]: ", sc.bos,index)
	flag, board := sc.bos[index-1].Guess()
	if(flag){
		sc.appendBoard(board)
	}
	return flag
}

func (sc StepContainer) Before(){
	flag := len(sc.bos)-1
	sc.bos = sc.bos[:len(sc.bos)-1]
	fmt.Println(sc)
	for ;sc.bos[flag].IsLast();{
		flag--
		sc.bos = sc.bos[:len(sc.bos)-1]
	}
}
