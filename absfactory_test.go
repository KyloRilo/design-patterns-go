package main

import (
	"fmt"
	"testing"
)

type IMerch interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type IShoe IMerch
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

type IShirt IMerch
type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

const (
	ADIDAS = "adidas"
	NIKE   = "nike"
)

type Adidas struct{}
type AdidasShoe struct {
	Shoe
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: ADIDAS,
			size: 14,
		},
	}
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: ADIDAS,
			size: 14,
		},
	}
}

type Nike struct{}
type NikeShoe struct {
	Shoe
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: NIKE,
			size: 14,
		},
	}
}

type NikeShirt struct {
	Shirt
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: NIKE,
			size: 14,
		},
	}
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case ADIDAS:
		return &Adidas{}, nil
	case NIKE:
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

func TestAbsFactory(t *testing.T) {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
