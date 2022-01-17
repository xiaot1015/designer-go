package main

import (
	"bridge/app/service/computer"
	"bridge/app/service/printer"
)

func main(){

	hp := &printer.Hp{}
	epson := &printer.Epson{}

	mac := computer.Mac{}
	win := computer.Windows{}

	mac.SetPrinter(hp)
	mac.Print()

	win.SetPrinter(epson)
	win.Print()
}