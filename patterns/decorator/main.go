package decorator

type IPizza interface {
	GetPrice() int
}

type VeggieMania struct{}

func (p *VeggieMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 7
}

type CheeseTopping struct {
	Pizza IPizza
}

func (t *CheeseTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 10
}
