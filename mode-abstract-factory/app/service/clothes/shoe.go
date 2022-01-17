package clothes

type Shoe struct {
	Size int
	Logo string
}

func(s *Shoe)SetSize(size int){
	s.Size = size
}
func(s *Shoe)GetSize() int{
	return s.Size
}
func(s *Shoe)SetLogo(logo string){
	s.Logo = logo
}
func(s *Shoe)GetLogo() string{
	return s.Logo
}
