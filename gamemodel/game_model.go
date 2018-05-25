package gamemodel

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/resources"
	"log"
	"os"
	"github.com/xosmig/roguelike/core/objects/factory"
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

type ExitHandler interface {
	OnExit()
}

type exit struct {
	objects.PositionData
	model *gameModel
}

func (e *exit) Interact(other objects.GameObject) {}
func (e *exit) Response(other objects.GameObject) {
	e.model.eventHandler.OnExit()
	// self-remove
	e.model.GetMap().Get(e.GetPosition()).Object = objects.Empty
}

func (e *exit) ModelName() string {
	return "exit"
}

func New(loader resources.Loader, mapName string, eventHandler EventHandler) (GameModel, error) {
	model := new(gameModel)

	exitObj := &exit{model: model}
	char := &character.Character{
		UnitData: objects.UnitData{
			MaxHP:    100,
			CurHP:    100,
			Team:     0,
		},
	}
	levelMap, err := gamemap.Load(loader, mapName, map[byte]factory.ObjectFactory{
		'#': factory.Repeated(objects.Wall),
		'O': factory.Singleton(exitObj),
		'@': factory.Singleton(char),
		'*': factory.Repeated(objects.Wall), // TODO
	})

	if err != nil {
		return nil, err
	}

	model.levelMap = levelMap
	model.char = char
	model.units = nil  // TODO
	model.eventHandler = eventHandler
	model.logger = log.New(os.Stderr, "", 0)

	return model, nil
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
