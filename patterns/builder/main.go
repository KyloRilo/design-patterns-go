package builder

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type Director struct {
	builder IBuilder
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

const (
	NORM  = "normal"
	IGLOO = "igloo"
)

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (b *NormalBuilder) setWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.Floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (b *IglooBuilder) setWindowType() {
	b.WindowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.DoorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.Floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case NORM:
		return newNormalBuilder()
	case IGLOO:
		return newIglooBuilder()
	}

	return nil
}
