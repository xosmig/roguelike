package gamemodel

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/resources"
	"log"
	"os"
)

type GameModel interface {
	GetMap() gamemap.GameMap
	GetCharacter() *character.Character
	DoMove(objects.Direction)
	SetLogger(*log.Logger)
}

type gameModel struct {
	levelMap     gamemap.GameMap
	char         *character.Character
	units        []objects.Unit
	eventHandler EventHandler
	logger       *log.Logger
}

func New(loader resources.Loader, mapName string, eventHandler EventHandler) (GameModel, error) {
	exitObj := objects.NewExit(eventHandler)
	char := &character.Character{
		UnitData: objects.UnitData{
			MaxHP:    100,
			CurHP:    100,
			Team:     0,
			Position: objects.Location{},
		},
	}
	levelMap, err := gamemap.Load(loader, mapName, map[byte]objects.GameObject{
		'#': objects.Wall,
		'O': exitObj,
		'@': char,
		'*': objects.Wall, // TODO
	})

	for row := 0; row < levelMap.GetHeight(); row++ {
		for col := 0; col < levelMap.GetWidth(); col++ {
			cell := levelMap.Get(objects.Loc(row, col))
			if unit, ok := cell.Object.(objects.Unit); ok {
				unit.SetPosition(objects.Loc(row, col))
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return &gameModel{
		levelMap: levelMap,
		char: char,
		units: nil,
		eventHandler: eventHandler,
		logger: log.New(os.Stderr, "", 0),
	}, nil
}

func (m *gameModel) SetLogger(logger *log.Logger) {
	m.logger = logger
}

func (m *gameModel) tryMove(unit objects.Unit, direction objects.Direction) {
	pos := unit.GetPosition()
	newPos := pos.Next(direction)

	m.logger.Printf("Debug: move %T %v from %v to %v\n", unit, direction, pos, newPos)

	oldCell := m.levelMap.Get(pos)
	newCell := m.levelMap.Get(newPos)

	unit.Interact(newCell.Object)
	newCell.Object.Response(unit)

	if newCell.Object != objects.Empty {
		m.logger.Printf("Debug: can't go because of %T\n", newCell.Object)
		return
	}

	unit.SetPosition(newPos)
	oldCell.Object = objects.Empty
	newCell.Object = unit
}

func (m *gameModel) DoMove(direction objects.Direction) {
	m.tryMove(m.GetCharacter(), direction)
}

func (m *gameModel) GetMap() gamemap.GameMap {
	return m.levelMap
}

func (m *gameModel) GetCharacter() *character.Character {
	return m.char
}
