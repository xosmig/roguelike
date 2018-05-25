package gamemodel

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/resources"
	"log"
	"github.com/xosmig/roguelike/core/objects/factory"
	"github.com/xosmig/roguelike/core/state"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/enemies/zombie"
)

type GameModel interface {
	state.GameState
	DoMove(geom.Direction)
}

type gameModel struct {
	levelMap     gamemap.GameMap
	char         *character.Character
	units        []unit.Unit
	eventHandler EventHandler
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
	char := character.New()
	levelMap, err := gamemap.Load(loader, mapName, map[byte]factory.ObjectFactory{
		'#': factory.Repeated(objects.Wall),
		'O': factory.Singleton(exitObj),
		'@': factory.Singleton(char),
		'z': factory.Singleton(zombie.New()),
		'$': factory.Singleton(objects.Wall), // TODO
	})

	if err != nil {
		return nil, err
	}

	model.levelMap = levelMap
	model.char = char
	model.units = nil // TODO
	model.eventHandler = eventHandler

	return model, nil
}

func (m *gameModel) TryMove(unit unit.Unit, direction geom.Direction) {
	pos := unit.GetPosition()
	newPos := pos.Next(direction)

	log.Printf("Debug: move %T %v from %v to %v\n", unit, direction, pos, newPos)

	oldCell := m.levelMap.Get(pos)
	newCell := m.levelMap.Get(newPos)

	unit.Interact(newCell.Object)
	newCell.Object.Response(unit)

	if newCell.Object != objects.Empty {
		log.Printf("Debug: %T can't go because of %T\n", unit, newCell.Object)
		return
	}

	unit.SetPosition(newPos)
	oldCell.Object = objects.Empty
	newCell.Object = unit
}

func (m *gameModel) DoMove(direction geom.Direction) {
	m.TryMove(m.GetCharacter(), direction)
}

func (m *gameModel) GetMap() gamemap.GameMap {
	return m.levelMap
}

func (m *gameModel) GetCharacter() *character.Character {
	return m.char
}
