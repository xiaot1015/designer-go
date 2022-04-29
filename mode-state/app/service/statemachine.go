package service

import "state/app/inter"

type StateMachine struct {
	hasItem inter.State
	itemRequested inter.State
	hasMoney inter.State
	noItem inter.State

	CurrentState inter.State

	itemCount int
	itemPrice int
}

func NewStateMachine() *StateMachine {
	obj := new(StateMachine)
	obj.itemCount = 100
	obj.itemPrice = 1000

	hasItemState := NewHasItem(obj)
	noItemState := NewNoItem(obj)
	hasMoneyState := NewHasMoney(obj)
	itemRequestedState := NewItemRequested(obj)

	obj.hasItem = hasItemState
	obj.noItem = noItemState
	obj.hasMoney = hasMoneyState
	obj.itemRequested = itemRequestedState
	obj.SetState(hasItemState)
	return obj
}

func (s *StateMachine) SetState(state inter.State) {
	s.CurrentState = state
}

func (s *StateMachine) IncrStore() {
	s.itemCount += 1
}

func (s *StateMachine) AddItem () {
	s.CurrentState.AddItem()
	return
}

func (s *StateMachine) InsertMoney(){
	s.CurrentState.InsertMoney()
	return
}

func (s *StateMachine) RequestItem(){
	s.CurrentState.RequestItem()
	return
}

func (s *StateMachine) DispenseItem() {
	s.CurrentState.DispenseItem()
	return
}