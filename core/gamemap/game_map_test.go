package gamemap

import (
	"testing"
	"bytes"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/objects/mock_objects"
	"github.com/golang/mock/gomock"
	"strings"
	"github.com/xosmig/roguelike/resources/mock_resources"
)

func TestLoad(t *testing.T) {
	ctrl := gomock.NewController(t)

	data := "6 11\n" +
		"###########\n" +
		"#@.#*.##.O#\n" +
		"#.#..###.##\n" +
		"#..#..##..#\n" +
		"#.........#\n" +
		"###########\n" +
		""
	height := 6
	width := 11

	loader := mock_resources.NewMockLoader(ctrl)
	loader.EXPECT().Load("maps/example").Times(1).Return(bytes.NewBufferString(data), nil)

	mapping := map[byte]objects.GameObject{
		'#': objects.Wall,
		'@': mock_objects.NewMockGameObject(ctrl),
		'*': mock_objects.NewMockGameObject(ctrl),
		'O': mock_objects.NewMockGameObject(ctrl),
	}

	gameMap, err := Load(loader, "maps/example", mapping)
	if err != nil {
		t.Fatalf("Error loading example map: %v", err)
	}

	lines := strings.Split(data, "\n")[1:]
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if mapping[lines[row][col]] != gameMap.Get(objects.Location{row, col}) {
				t.Errorf("Invalid object at position (%d, %d)", row, col)
			}
		}
	}
}
