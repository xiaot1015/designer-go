package main

import (
	"mediator/app/service"
	"mediator/app/service/train"
)

func main(){
	stationManager := service.NewStation(true)

	passengerTrain := train.NewPassengerTrain(stationManager)
	freightTrain := train.NewFreightTrain(stationManager)

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
	passengerTrain.Depart()

}