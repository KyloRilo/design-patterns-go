package flyweight

import "fmt"

const (
	TPlayerType = iota
	CTPlayerType
)

const (
	TDressType = iota
	CTDressType
)

type dress interface {
	GetColor() string
}

type TDress struct {
	color string
}

func (d *TDress) GetColor() string {
	return d.color
}

func newTDress() *TDress {
	return &TDress{color: "red"}
}

type CTDress struct {
	color string
}

func (d *CTDress) GetColor() string {
	return d.color
}

func newCTDress() *CTDress {
	return &CTDress{color: "green"}
}

type DressFactory struct {
	DressMap map[int]dress
}

func (d *DressFactory) getDressByType(dressType int) (dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TDressType {
		d.DressMap[dressType] = newTDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CTDressType {
		d.DressMap[dressType] = newCTDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

var dressFactorySingleton = &DressFactory{
	DressMap: make(map[int]dress),
}

func GetDressFactorySingleton() *DressFactory {
	return dressFactorySingleton
}

type Player struct {
	dress      dress
	playerType int
	lat        int
	long       int
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

func newPlayer(playerType, dressType int) *Player {
	dress, _ := GetDressFactorySingleton().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func (g *game) AddTerrorist() {
	g.terrorists = append(g.terrorists, newPlayer(TPlayerType, TDressType))
}

func (g *game) AddCounterTerrorist() {
	g.counterTerrorists = append(g.counterTerrorists, newPlayer(CTPlayerType, CTDressType))
}

func NewGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}
