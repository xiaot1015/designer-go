package service


type HasMoney struct {
	stateMachine *StateMachine
}

func NewHasMoney(state *StateMachine) *HasMoney {
	obj := new(HasMoney)
	obj.stateMachine = state
	return obj
}

func (n *HasMoney) AddItem () {

}

func (n *HasMoney) InsertMoney(){

}

func (n *HasMoney) RequestItem(){

}

func (n *HasMoney) DispenseItem() {

}