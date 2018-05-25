package geom

import (
	"fmt"
	"github.com/xosmig/roguelike/util"
)

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

func (loc Location) Neighbour(other Location) bool {
	return (util.AbsInt(loc.Row-other.Row) + util.AbsInt(loc.Col-other.Col)) == 1
}

func (loc Location) StepTo(other Location) (Direction, bool) {
	switch loc.Row - other.Row {
	case 0:
	case -1:
		return Up, true
	case 1:
		return Down, true
	default:
		return Nowhere, false
	}

	switch loc.Col - other.Col {
	case 0:
	case -1:
		return Left, true
	case 1:
		return Right, true
	default:
		return Nowhere, false
	}

	// same location
	return Nowhere, true
}
