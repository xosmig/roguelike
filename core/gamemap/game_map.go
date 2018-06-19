package gamemap

import (
	"fmt"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/objects/factory"
	"github.com/xosmig/roguelike/resources"
	"unicode"
)

// Cell represents a single position on a game map.
type Cell struct {
	Object objects.GameObject
}

// GameMap represents a rectangular map.
type GameMap interface {
	GetHeight() int
	GetWidth() int
	Get(geom.Location) *Cell
}

// AllObjects returns a list of all objects on the map.
func AllObjects(gameMap GameMap) []objects.GameObject {
	res := make([]objects.GameObject, 0, gameMap.GetHeight()*gameMap.GetWidth())
	for row := 0; row < gameMap.GetHeight(); row++ {
		for col := 0; col < gameMap.GetWidth(); col++ {
			res = append(res, gameMap.Get(geom.Loc(row, col)).Object)
		}
	}
	return res
}

// Remove replaces the object at the position pos with objects.Empty.
func Remove(gameMap GameMap, pos geom.Location) {
	gameMap.Get(pos).Object = objects.Empty
}

type staticMap struct {
	cells [][]Cell
}

func (m *staticMap) GetHeight() int {
	return len(m.cells)
}

func (m *staticMap) GetWidth() int {
	return len(m.cells[0])
}

func (m *staticMap) Get(loc geom.Location) *Cell {
	return &m.cells[loc.Row][loc.Col]
}

// Load tries to load a static game map, described in the given resource within the given loader.
//
// mapping defines how to load the map:
// The object returned by mapping[ch].Create(geom.Loc(x, y)) will be placed at the position (x, y),
// where ch is the char at the position (x, y) in the input.
//
// If mapping['.'] is not defined, it will be assigned to objects.Empty.
func Load(loader resources.Loader, name string, mapping map[byte]factory.ObjectFactory) (GameMap, error) {
	reader, err := loader.Load(name)
	if err != nil {
		return nil, err
	}

	if _, present := mapping['.']; !present {
		mapping['.'] = factory.Repeated(objects.Empty)
	}

	var height int
	var width int
	if _, err = fmt.Fscanf(reader, "%d %d\n", &height, &width); err != nil {
		return nil, fmt.Errorf("while reading map header: %v", err)
	}

	if height == 0 || width == 0 {
		return nil, fmt.Errorf("invalid map size: %d rows and %d columns", height, width)
	}

	cells := make([][]Cell, height)
	for row := range cells {
		cells[row] = make([]Cell, width)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			value := ' '
			for unicode.IsSpace(value) {
				if _, err = fmt.Fscanf(reader, "%c", &value); err != nil {
					return nil, fmt.Errorf("while reading map position (%d, %d): %v", row, col, err)
				}
			}
			f, present := mapping[byte(value)]
			if !present {
				return nil, fmt.Errorf("no mapping for '%c'", value)
			}
			obj, err := f.Create(geom.Loc(row, col))
			if err != nil {
				return nil, fmt.Errorf("while create game object: %v", err)
			}
			cells[row][col] = Cell{Object: obj}
		}
	}

	return &staticMap{cells}, nil
}
