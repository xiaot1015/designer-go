package computer

import (
	"bridge/app/inter"
	"fmt"
)

type Mac struct {

	pinter inter.Printer
}

func (m *Mac)Print(){
	fmt.Println("mac print")
	m.pinter.PrintFile()
}

func (m *Mac)SetPrinter(printer inter.Printer){
	m.pinter = printer
}

