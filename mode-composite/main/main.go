package main

import (
	"composite/app/service"
	"fmt"
)

// 组件模式
func main(){
	p1 := service.NewProduct(3)
	p2 := service.NewProduct(5)

	pp1 := service.NewPromotion(11)
	pp2 := service.NewPromotion(22)

	o1 := service.NewOrder(1123)
	o1.SetComponents(p1,p2,pp1,pp2)
	if err := o1.CheckValid(); err != nil {
		fmt.Println("o1 checkvalid:",err)
	}
	o2 := service.NewOrder(3311)
	o2.SetComponents(p1,pp1)
	if err := o2.CheckValid(); err != nil {
		fmt.Println("o2 checkvalid:",err)
	}

}