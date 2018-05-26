package gamemap

import (
	"fmt"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/objects/factory"
	"github.com/xosmig/roguelike/resources"
	"unicode"
)

type Cell struct {
	Object objects.GameObject
}

type GameMap interface {
	GetHeight() int
	GetWidth() int
	Get(geom.Location) *Cell
}

func AllObjects(gameMap GameMap) []objects.GameObject {
	res := make([]objects.GameObject, 0, gameMap.GetHeight()*gameMap.GetWidth())
	for row := 0; row < gameMap.GetHeight(); row++ {
		for col := 0; col < gameMap.GetWidth(); col++ {
			res = append(res, gameMap.Get(geom.Loc(row, col)).Object)
		}
	}
	return res
}

func Remove(gameMap GameMap, pos geom.Location) {
	gameMap.Get(pos).Object = objects.Empty
}

type StaticMap struct {
	cells [][]Cell
}

func (m *StaticMap) Get(loc geom.Location) *Cell {
	return &m.cells[loc.Row][loc.Col]
}

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

	return &StaticMap{cells}, nil
}

func (m *StaticMap) GetHeight() int {
	return len(m.cells)
}

func (m *StaticMap) GetWidth() int {
	return len(m.cells[0])
}
