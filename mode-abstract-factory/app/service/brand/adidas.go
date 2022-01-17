package brand

import (
	"abstract-factory/app/inter"
	"abstract-factory/app/service/clothes"
	"abstract-factory/app/service/clothes/shoe"
)

type Adidas struct {

}

func (a *Adidas)MakeShoe() inter.IShoe {
	return  &shoe.AdidasShoe{
		clothes.Shoe{20, "Adidas shoe"}}
}

func (a *Adidas)MakeShirt() inter.IShirt{
	return  &shoe.AdidasShirt{
		clothes.Shirt{"blue", "Adidas shirt"}}
}