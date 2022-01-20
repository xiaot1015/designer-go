package service

import "mediator/app/inter"

type Station struct {
	IsPlatformFree bool
	TrainQueue     []inter.ITrain
}

func NewStation(isFree bool) *Station{
	obj := new(Station)
	obj.IsPlatformFree = isFree
	return obj
}

func (s *Station)CanArrive(train inter.ITrain) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, train)
	return false
}

func (s *Station)NotifyAboutDeparture() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrain := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrain.PreArrival()
	}
}