package main

import (
	"abstract-factory/app/factory"
	"abstract-factory/app/inter"
	"fmt"
)


func main() {
	adidasFactory := factory.SportsManager(factory.Brand_Adidas)
	nikeFactory := factory.SportsManager(factory.Brand_Nike)

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s inter.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s inter.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %s", s.GetColor())
	fmt.Println()
}