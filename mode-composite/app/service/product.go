package service

import (
	"errors"
	"fmt"
)

type Product struct {
	ProductId int
}
func NewProduct(ProductId int) *Product{
	obj := new(Product)
	obj.ProductId = ProductId
	return obj
}
func (p *Product)CheckValid() error {
	fmt.Println("product CheckValid", p.ProductId)
	if p.ProductId % 2 == 0 {
		 return errors.New("not valid")
	}
	return nil
}

func (p *Product)SetProductId(productId int){
	p.ProductId = productId
}
func (p *Product)GetProductId() int{
	return p.ProductId
}