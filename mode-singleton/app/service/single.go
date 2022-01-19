package service

import (
	"fmt"
	"sync"
)

var once sync.Once
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}


func GetInstanceV2() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Create single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}