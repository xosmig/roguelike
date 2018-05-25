package character

import (
	"github.com/xosmig/roguelike/core/objects"
)

type Character struct {
	objects.UnitData
	Wearing   Item
	Inventory []Item
}

func (char *Character) RecvDamage(dmg int, from objects.Unit) {
	objects.RecvDamageDefault(char, dmg, from)
}

func (char *Character) Die(from objects.Unit) {
	// TODO
}

func (char *Character) Interact(other objects.GameObject) {
	objects.UnitInteractDefault(char, other)
}

func (char *Character) Response(other objects.GameObject) {
	char.Interact(other)
}

func (char *Character) ModelName() string {
	return "character"
}
