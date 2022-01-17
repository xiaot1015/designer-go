package inter

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
