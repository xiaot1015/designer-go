package train

import (
	"fmt"
	"mediator/app/inter"
)

type PassengerTrain struct {
	Station inter.IStation
}

func NewPassengerTrain(station inter.IStation) *PassengerTrain{
	obj := new(PassengerTrain)
	obj.Station = station
	return obj
}

func (p *PassengerTrain)Arrive(){
	if !p.Station.CanArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain)PreArrival(){
	fmt.Println("PassengerTrain: PreArrival")
	p.Arrive()
}

func (p *PassengerTrain)Depart(){
	fmt.Println("PassengerTrain: Depart")
	p.Station.NotifyAboutDeparture()
}