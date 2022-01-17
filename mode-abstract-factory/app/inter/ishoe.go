package inter

type IShoe interface {
	SetLogo(logo string)
	GetLogo() string
	SetSize(size int)
	GetSize() int
}