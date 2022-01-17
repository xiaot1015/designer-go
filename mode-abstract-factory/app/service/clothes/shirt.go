package clothes

type Shirt struct {
	Color string
	Logo string
}

func(s *Shirt)SetColor(color string){
	s.Color = color
}
func(s *Shirt)GetColor() string{
	return s.Color
}
func(s *Shirt)SetLogo(logo string){
	s.Logo = logo
}
func(s *Shirt)GetLogo() string{
	return s.Logo
}