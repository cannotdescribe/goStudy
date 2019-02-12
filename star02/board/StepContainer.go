package board

import (
	"go_code/star02/utils"
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

func (sc *StepContainer) appendCopyBoard(){
	index := len(*sc.bos)
	lastBoard := (*sc.bos)[index-1]

	copyBoard := lastBoard.CopyBoard()

	flag := lastBoard.flag+1
	for;lastBoard.stepContainer.isInitData(flag);{
		flag++
	}
	copyBoard.flag = flag

	sc.appendBoard(copyBoard)
}

func (sc *StepContainer) NextStep() bool{
	index := len(*sc.bos)
	
	flag := (*sc.bos)[index-1].Guess()
	if(flag){
		sc.appendCopyBoard()
	}
	return flag
}

func (sc *StepContainer) Before() bool{
	lastEndFlag, endWatch := (*sc.bos)[len(*sc.bos)-1].LastSecondValue(sc.flag)
	*sc.bos = (*sc.bos)[:len(*sc.bos)-1]
	flag := len(*sc.bos)-1
	if(lastEndFlag == -1){
		return false
	}
	for ;flag>0 && (*sc.bos)[flag].IsLast(endWatch); {
		flag--
		lastEndFlag, endWatch = (*sc.bos)[len(*sc.bos)-1].LastSecondValue(sc.flag)
		if(lastEndFlag == -1){
			return false
		}
		*sc.bos = (*sc.bos)[:len(*sc.bos)-1]
	}

	return true
}

func (sc *StepContainer) EndBoard() []int{
	return (*sc.bos)[len(*sc.bos)-1].arr
}

func (sc *StepContainer) End() bool{
	flag := 81-sc.flag.Size()
	return len(*sc.bos) == flag+1
}