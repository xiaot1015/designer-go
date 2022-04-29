package main

import "state/app/service"

func main(){
	machine := service.NewStateMachine()
	machine.RequestItem()

	machine.InsertMoney()

	machine.DispenseItem()

	machine.AddItem()



	machine.RequestItem()

	machine.InsertMoney()

	machine.DispenseItem()

}