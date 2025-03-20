package main

import (
	"fmt"
	"sync"
	"testing"
)

var lock = &sync.Mutex{}

type single struct{}

var singleton *single

func getInstance() *single {
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

func TestSingleton(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
