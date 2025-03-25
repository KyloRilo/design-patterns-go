package tests

import (
	"fmt"
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/factory"
)

func TestFactory(t *testing.T) {
	ak47, _ := factory.GetGun(factory.AK47)
	musket, _ := factory.GetGun(factory.MUSKET)

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
