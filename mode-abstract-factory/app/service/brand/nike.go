package brand

import (
	"abstract-factory/app/inter"
	"abstract-factory/app/service/clothes"
	"abstract-factory/app/service/clothes/shoe"
)

type Nike struct {

}

func (n *Nike)MakeShoe() inter.IShoe {
	return  &shoe.NikeShoe{
		clothes.Shoe{10, "nick shoe"}}
}

func (n *Nike)MakeShirt() inter.IShirt{
	return  &shoe.NikeShirt{
		clothes.Shirt{"red", "nick shirt"}}

}