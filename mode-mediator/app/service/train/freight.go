package train

import (
	"fmt"
	"mediator/app/inter"
)

type FreightTrain struct {
	Station inter.IStation
}

func NewFreightTrain(station inter.IStation) *FreightTrain{
	obj := new(FreightTrain)
	obj.Station = station
	return obj
}

func (f *FreightTrain)Arrive(){
	if !f.Station.CanArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f *FreightTrain)PreArrival(){
	fmt.Println("FreightTrain: PreArrival")
	f.Arrive()
}

func (f *FreightTrain)Depart(){
	fmt.Println("FreightTrain: Depart")
	f.Station.NotifyAboutDeparture()
}