package service

import (
	"composite/app/inter"
	"errors"
	"fmt"
)

type Order struct {
	OId int
	PromotionId int
	ProductId int
	Components []inter.Component
}

func NewOrder(oid int) *Order{
	obj := new(Order)
	obj.OId = oid
	return obj
}

func (o *Order)CheckValid() error {
	fmt.Println("Order CheckValid", o.OId)
	if o.OId % 2 == 0 {
		return errors.New("order not valid")
	}
	if len(o.Components) < 1 {
		return nil
	}
	for i:=0; i<len(o.Components); i++ {
		component := o.Components[i]
		if err := component.CheckValid(); err != nil {
			return err
		}
	}
	return nil
}

func (o *Order)SetOrderId(oid int){
	o.OId = oid
}

func (o *Order)GetOrderId() int{
	return o.OId
}

func (o *Order)SetComponents(component... inter.Component) {
	o.Components = append(o.Components, component...)
}
func (o *Order)GetComponents() []inter.Component {
	return o.Components
}