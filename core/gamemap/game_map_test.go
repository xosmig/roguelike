package gamemap

import (
	"testing"
	"bytes"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/golang/mock/gomock"
	"strings"
	"github.com/xosmig/roguelike/resources/mock_resources"
	"github.com/xosmig/roguelike/core/objects/factory"
	"github.com/xosmig/roguelike/core/objects/factory/mock_factory"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/geom"
)

func TestLoad(t *testing.T) {
	ctrl := gomock.NewController(t)

	data := "6 11\n" +
		"###########\n" +
		"#..#..##..#\n" +
		"#@#..###.##\n" +
		"#..#..##..#\n" +
		"#.........#\n" +
		"###########\n" +
		""
	height := 6
	width := 11

	loader := mock_resources.NewMockLoader(ctrl)
	loader.EXPECT().Load("maps/example").Times(1).Return(bytes.NewBufferString(data), nil)

	wallFactory := mock_factory.NewMockObjectFactory(ctrl)
	wallFactory.EXPECT().Create(gomock.Any()).Times(42).Return(objects.Wall, nil)

	char := &character.Character{}
	charFactory := mock_factory.NewMockObjectFactory(ctrl)
	charFactory.EXPECT().Create(geom.Loc(2, 1)).Times(1).Return(char, nil)

	mapping := map[byte]factory.ObjectFactory{
		'#': wallFactory,
		'@': charFactory,
	}

	gameMap, err := Load(loader, "maps/example", mapping)
	if err != nil {
		t.Fatalf("Error loading example map: %v", err)
	}

	objMapping := map[byte]objects.GameObject{
		'#': objects.Wall,
		'@': char,
		'.': objects.Empty,
	}

	lines := strings.Split(data, "\n")[1:]
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			loc := geom.Loc(row, col)
			if objMapping[lines[row][col]] != gameMap.Get(loc).Object {
				t.Errorf("Wrong object at %v", loc)
			}
		}
	}

}
