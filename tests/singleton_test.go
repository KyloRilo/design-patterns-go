package tests

import (
	"fmt"
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/singleton"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()
}
