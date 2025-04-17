package tests

import (
	"fmt"
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/flyweight"
)

func TestFlyweight(t *testing.T) {
	game := flyweight.NewGame()

	//Add Terrorist
	game.AddTerrorist()
	game.AddTerrorist()
	game.AddTerrorist()
	game.AddTerrorist()

	//Add CounterTerrorist
	game.AddCounterTerrorist()
	game.AddCounterTerrorist()
	game.AddCounterTerrorist()

	dressFactoryInstance := flyweight.GetDressFactorySingleton()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
