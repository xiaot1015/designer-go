package service


type HasItem struct {
	stateMachine *StateMachine
}

func NewHasItem(state *StateMachine) *HasItem {
	obj := new(HasItem)
	obj.stateMachine = state
	return obj
}

func (s *HasItem) AddItem () {

}
func (s *HasItem) InsertMoney(){

}
func (s *HasItem) RequestItem(){

}
func (s *HasItem) DispenseItem() {

}