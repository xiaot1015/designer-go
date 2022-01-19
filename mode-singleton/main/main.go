package main

import (
	"fmt"
	"singleton/app/service"
)

func main() {

	//for i := 0; i < 30; i++ {
	//	go service.GetInstance()
	//}
	//
	//time.Sleep(3*time.Second)

	for i := 0; i < 30; i++ {
		go service.GetInstanceV2()
	}
	
	fmt.Scanln()
}