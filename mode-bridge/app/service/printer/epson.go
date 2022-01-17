package printer

import "fmt"

type Epson struct {

}

func (E *Epson)PrintFile(){
	fmt.Println("Epson print:")
}