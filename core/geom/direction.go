package geom

import "fmt"

type Direction int

const (
	Nowhere = iota
	Up
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	default:
		panic(fmt.Sprintf("Unknown direction id: %d", d))
	}
}
