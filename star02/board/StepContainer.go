package board

import (
	"go_code/star02/utils"
)

type StepContainer struct{
	bos []Board
	flag *utils.Set
}

func NewStepContainer(sodu [81]int) *StepContainer{
	stepContainer := &StepContainer{
		bos: make([]Board, 0),
		flag: utils.NewSet(map[interface{}]bool{}),
	}
	// stepContainer.flag = utils.NewSet(map[interface{}]bool{})


	// stepContainer.bos[0] = NewBoard(sodu, 0)
	for index, value := range sodu{
		if(value != 0){
			stepContainer.flag.Add(index)
		}
	}
	return stepContainer
}

func (sc StepContainer) CanNextStep() bool{
	index := len(sc.bos)
	return sc.bos[index-1].Check(1)
}

func (sc StepContainer) NextStep(){
	// index := len(sc.bos)
	// sc.bos[index]
}

func (sc StepContainer) Before(){
	
}
