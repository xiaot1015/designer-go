package service


type ItemRequested struct {
	stateMachine *StateMachine
}

func NewItemRequested(state *StateMachine) *ItemRequested {
	obj := new(ItemRequested)
	obj.stateMachine = state
	return obj
}

func (n *ItemRequested) AddItem () {

}
func (n *ItemRequested) InsertMoney(){

}
func (n *ItemRequested) RequestItem(){

}
func (n *ItemRequested) DispenseItem() {

}