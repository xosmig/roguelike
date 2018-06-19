package gamemodel

import (
	"github.com/xosmig/roguelike/core/ai"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/enemies/zombie"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/items"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/objects/factory"
	"github.com/xosmig/roguelike/core/state"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/gamemodel/status"
	"github.com/xosmig/roguelike/resources"
	"log"
	"math/rand"
	"time"
)

// GameModel manipulates game logic. Supposed to be called from the user interface implementation.
type GameModel interface {
	state.GameState
	DoMove(geom.Direction)
	Status() status.Status
}

type gameModel struct {
	levelMap gamemap.GameMap
	char     character.Character
	status   status.Status
}

type exit struct {
	objects.PositionData
	model *gameModel
}

func (e *exit) Response(other objects.GameObject) {
	if _, ok := other.(character.Character); !ok {
		return
	}

	log.Println("Info: Character reached exit. Victory.")
	e.model.status = status.Victory
	gamemap.Remove(e.model.GetMap(), e.GetPosition())
}

func (e *exit) ModelName() string {
	return "exit"
}

// New initializes a new `GameModel` with the given map.
// Returns error when the map cannot be loaded.
func New(loader resources.Loader, mapName string) (GameModel, error) {
	model := new(gameModel)
	rand.Seed(time.Now().UnixNano())

	exitObj := &exit{model: model}
	char := character.New()
	levelMap, err := gamemap.Load(loader, mapName, map[byte]factory.ObjectFactory{
		'#': factory.Repeated(objects.Wall),
		'O': factory.Singleton(exitObj),
		'@': factory.Singleton(char),
		'z': factory.Singleton(zombie.New()),
		'$': factory.Singleton(items.NewItemObject(model, items.NewHealthAmulet())),
	})

	if err != nil {
		return nil, err
	}

	model.levelMap = levelMap
	model.char = char
	model.status = status.Continue

	return model, nil
}

// See `GameState` documentation
func (m *gameModel) TryMove(obj objects.MovableObject, direction geom.Direction) {
	pos := obj.GetPosition()
	newPos := pos.Next(direction)
	if pos == newPos {
		log.Printf("Debug: %T stays at %v\n", obj, pos)
		return
	}

	log.Printf("Debug: move %T %v from %v to %v\n", obj, direction, pos, newPos)

	oldCell := m.levelMap.Get(pos)
	newCell := m.levelMap.Get(newPos)

	obj.Interact(newCell.Object)
	newCell.Object.Response(obj)

	if newCell.Object != objects.Empty {
		log.Printf("Debug: %T can't go because of %T\n", obj, newCell.Object)
		return
	}

	obj.SetPosition(newPos)
	oldCell.Object = objects.Empty
	newCell.Object = obj
}

// DoMove tries to move the character to the given direction (see `TryMove`).
// After that all other objects do their actions.
// At the end of the turn all dead object are removed from map.
func (m *gameModel) DoMove(direction geom.Direction) {
	if m.Status() != status.Continue {
		log.Println("Error: trying to move, when game already finished. Might be some bug in the UI.")
		return
	}

	allObjects := gamemap.AllObjects(m.levelMap)

	m.TryMove(m.GetCharacter(), direction)
	for _, obj := range allObjects {
		if actionable, ok := obj.(ai.Actionable); ok {
			log.Printf("Debug: %T doing action\n", obj)
			actionable.DoAction(m)
		}
	}

	for _, obj := range allObjects {
		if u, ok := obj.(unit.Unit); ok && !unit.IsAlive(u) {
			log.Printf("Info: removing dead %T from map\n", u)
			state.RemoveDead(m, u)
		}
	}

	if !unit.IsAlive(m.GetCharacter()) {
		log.Println("Info: Character died. Defeat.")
		m.status = status.Defeat
	}
}

func (m *gameModel) GetMap() gamemap.GameMap {
	return m.levelMap
}

func (m *gameModel) GetCharacter() character.Character {
	return m.char
}

func (m *gameModel) Status() status.Status {
	return m.status
}
