package service

import (
	"errors"
	"fmt"
)

type Promotion struct {
	PId int
}

func NewPromotion(pid int) *Promotion{
	obj := new(Promotion)
	obj.PId = pid
	return obj
}

func (p *Promotion)CheckValid() error {
	fmt.Println("Promotion CheckValid", p.PId)
	if p.PId % 2 == 0 {
		return errors.New("promotion not valid")
	}
	return nil
}

func (p *Promotion)SetPromotionId(pid int){
	p.PId = pid
}
func (p *Promotion)GetPromotionId() int{
	return p.PId
}