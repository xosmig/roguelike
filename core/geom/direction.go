package geom

import (
	"fmt"
	"math/rand"
)

// Direction a direction on game map
type Direction int

const (
	Nowhere Direction = iota
	Up
	Down
	Left
	Right
)

// RandomDirection returns random direction (out of 4). Doesn't return Nowhere
func RandomDirection() Direction {
	return Direction(rand.Intn(4) + int(Nowhere) + 1)
}

func (d Direction) String() string {
	switch d {
	case Nowhere:
		return "Nowhere"
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	default:
		return fmt.Sprintf("INVALID_DIRECTION(%d)", int(d))
	}
}
