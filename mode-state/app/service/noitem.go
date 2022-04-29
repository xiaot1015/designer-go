package service

import "fmt"

type NoItem struct {
	stateMachine *StateMachine
}
func NewNoItem(state *StateMachine) *NoItem {
	obj := new(NoItem)
	obj.stateMachine = state
	return obj
}
func (n *NoItem) AddItem () {
	n.stateMachine.IncrStore()
}

func (n *NoItem) InsertMoney(){
	fmt.Printf("out of stack")
}

func (n *NoItem) RequestItem(){
	fmt.Printf("out of stack")
}

func (n *NoItem) DispenseItem() {
	fmt.Printf("out of stack")

}