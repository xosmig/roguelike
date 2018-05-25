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

func AbsDist(a Location, b Location) int {
	return util.AbsInt(a.Row-b.Row) + util.AbsInt(a.Col-b.Col)
}

func Neighbours(a Location, b Location) bool {
	return AbsDist(a, b) == 1
}

func (loc Location) StepTo(other Location) (Direction, bool) {
	if AbsDist(loc, other) > 1 {
		return 0, false
	}

	switch other.Row - loc.Row {
	case -1:
		return Up, true
	case 1:
		return Down, true
	}

	switch other.Col - loc.Col {
	case -1:
		return Left, true
	case 1:
		return Right, true
	}

	return Nowhere, true
}
