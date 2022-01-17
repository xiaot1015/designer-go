package inter

type IShirt interface {
	SetLogo(logo string)
	GetLogo() string
	SetColor(color string)
	GetColor() string
}