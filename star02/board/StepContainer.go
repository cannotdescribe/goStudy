package board

import (
	"fmt"
	"go_code/star02/utils"
)

type StepContainer struct{
	bos []Board
	flag utils.Set
}

func newStepContainer(){
	return &StepContainer{
		make([]Board, 0),
		&utils.Set{
			m: map[interface{}]bool{},
		}
	}
}

func (sc StepContainer) Before(){
	fmt.Println("before");
}
