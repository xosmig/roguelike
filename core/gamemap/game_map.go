package gamemap

import (
	"github.com/xosmig/roguelike/resources"
	"fmt"
	"github.com/xosmig/roguelike/core/objects"
	"unicode"
)

type Cell struct {
	Object objects.GameObject
}

type GameMap interface {
	GetHeight() int
	GetWidth() int
	Get(objects.Location) *Cell
}

type StaticMap struct {
	cells [][]Cell
}

func (m *StaticMap) Get(loc objects.Location) *Cell {
	return &m.cells[loc.Row][loc.Col]
}

func Load(loader resources.Loader, name string, mapping map[byte]objects.GameObject) (GameMap, error) {
	reader, err := loader.Load(name)
	if err != nil {
		return nil, err
	}

	if _, present := mapping['.']; !present {
		mapping['.'] = objects.Empty
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
			obj, present := mapping[byte(value)]
			if !present {
				return nil, fmt.Errorf("no mapping for '%c'", value)
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
