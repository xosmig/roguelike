package character

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
)

type Character struct {
	unit.UnitData
	Wearing   Item
	Inventory []Item
}

func New() *Character {
	return &Character{
		UnitData: unit.UnitData{
			MaxHP: 3,
			CurHP: 3,
			Team:  unit.TeamGood,
		},
	}
}

func (char *Character) RecvDamage(dmg int, from unit.Unit) {
	unit.RecvDamageDefault(char, dmg, from)
}

func (char *Character) Die(from unit.Unit) {
	// TODO
}

func (char *Character) AddItem(item Item) {
	char.Inventory = append(char.Inventory, item)
}

func (char *Character) Interact(other objects.GameObject) {
	unit.InteractDefault(char, other)
}

func (char *Character) Response(other objects.GameObject) {
	// empty
}

func (char *Character) ModelName() string {
	return "character"
}
