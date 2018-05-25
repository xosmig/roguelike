package character

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
)

type Character struct {
	unit.UnitData
	wearing   Item
	inventory []Item
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

func (char *Character) Wearing() Item {
	return char.wearing
}

func (char *Character) Inventory() []Item {
	return char.inventory
}

func (char *Character) AddItem(item Item) {
	char.inventory = append(char.inventory, item)
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
