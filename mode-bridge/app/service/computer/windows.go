package computer

import (
	"bridge/app/inter"
	"fmt"
)

type Windows struct {
	printer inter.Printer
}

func (w *Windows)Print(){
	fmt.Println("Windows print")
	w.printer.PrintFile()
}

func (w *Windows)SetPrinter(printer inter.Printer){
	w.printer = printer
}
