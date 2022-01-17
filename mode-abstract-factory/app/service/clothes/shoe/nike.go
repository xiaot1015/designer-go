package shoe

import "abstract-factory/app/service/clothes"

type NikeShoe struct {
	clothes.Shoe
}

type NikeShirt struct {
	clothes.Shirt
}
