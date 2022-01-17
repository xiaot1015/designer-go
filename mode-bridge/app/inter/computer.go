package inter


type Computer interface {
	Print()
	SetPrinter(printer Printer)
}