package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleton *single

func GetInstance() *single {
	if singleton == nil {
		fmt.Println("Creating singleton now.")
		lock.Lock()
		defer lock.Unlock()
		singleton = &single{}
	} else {
		fmt.Println("Singleton already created.")
	}

	return singleton
}
