package main

import (
	"fmt"
	"testing"
)

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

const (
	AK47   = "AK47"
	MUSKET = "Musket"
)

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  AK47,
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  MUSKET,
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case AK47:
		return newAk47(), nil
	case MUSKET:
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

func TestFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
