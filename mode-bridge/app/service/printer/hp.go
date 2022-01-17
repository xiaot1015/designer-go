package printer

import "fmt"

type Hp struct {

}

func (h *Hp)PrintFile(){
	fmt.Println("hp print:")
}