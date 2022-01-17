package shoe

import "abstract-factory/app/service/clothes"

type AdidasShoe struct {
	clothes.Shoe
}

type AdidasShirt struct {
	clothes.Shirt
}