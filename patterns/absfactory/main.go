package absfactory

import (
	"fmt"
)

type IMerch interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type IShoe IMerch
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
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

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

const (
	ADIDAS = "adidas"
	NIKE   = "nike"
)

type Adidas struct{}
type AdidasShoe struct {
	Shoe
}

func (a *Adidas) MakeShoe() IShoe {
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

func (a *Adidas) MakeShirt() IShirt {
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

func (n *Nike) MakeShoe() IShoe {
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

func (n *Nike) MakeShirt() IShirt {
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
