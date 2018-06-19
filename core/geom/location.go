package geom

import (
	"fmt"
	"github.com/xosmig/roguelike/util"
)

// Location represents a point on game map
type Location struct {
	Row int
	Col int
}

// Loc returns new location with specified row and column
func Loc(row int, col int) Location {
	return Location{Row: row, Col: col}
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.Row, loc.Col)
}

// Next returns `loc` shifted by 1 position to the direction `direction`
func (loc Location) Next(direction Direction) Location {
	switch direction {
	case Nowhere:
		return loc
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

// ManhattanDist return manhattan distance between 2 locations
func ManhattanDist(a Location, b Location) int {
	return util.AbsInt(a.Row-b.Row) + util.AbsInt(a.Col-b.Col)
}

// StepTo returns boolean indicating whether the 2 locations are neighbours or equal.
// If they are, returns the direction from `loc` to `other` (`Nowhere` if they are equal).
func (loc Location) StepTo(other Location) (Direction, bool) {
	if ManhattanDist(loc, other) > 1 {
		return Nowhere, false
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
