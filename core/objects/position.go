package objects

import "fmt"

type Location struct {
	Row int
	Col int
}

func Loc(row int, col int) Location {
	return Location{Row: row, Col: col}
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.Row, loc.Col)
}

type Direction int

const (
	Up    = iota
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

func (loc Location) Next(direction Direction) Location {
	switch direction {
	case Up:
		return Location{loc.Row - 1, loc.Col}
	case Down:
		return Location{loc.Row + 1, loc.Col}
	case Left:
		return Location{loc.Row, loc.Col - 1}
	case Right:
		return Location{loc.Row, loc.Col + 1}
	default:
		panic(fmt.Sprint("Invalid direction id: ", direction))
	}
}
