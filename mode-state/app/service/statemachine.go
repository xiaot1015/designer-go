package service

import "state/app/inter"

type StateMachine struct {
	Ordered inter.State
	OrderPaying inter.State
	OrderPay inter.State
	OrderCancel inter.State
	OrderRefund inter.State

	CurrentState inter.State
}

func NewStateMachine() *StateMachine {
	obj := new(StateMachine)
	return obj
}