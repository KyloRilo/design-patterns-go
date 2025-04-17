package tests

import (
	"fmt"
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/decorator"
)

func TestDecorator(t *testing.T) {
	pizza := &decorator.VeggieMania{}

	pizzaCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	pizzaCheeseMato := &decorator.TomatoTopping{
		Pizza: pizzaCheese,
	}

	fmt.Printf("Price of VeggieMania pizza with extra cheese and tomato: %d", pizzaCheeseMato.GetPrice())
}
