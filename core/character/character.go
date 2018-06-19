package character

import (
	"fmt"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
)

type Character interface {
	unit.Unit
	Wearing() Item
	WearOrTakeOff(idx int) error
	AddItem(item Item)
	Inventory() []Item
}

type Item interface {
	Wear(character Character) error
	TakeOff(character Character) error
	IconName() string
}

type character struct {
	unit.UnitData
	wearing   Item
	inventory []Item
}

func New() Character {
	return &character{
		UnitData: unit.UnitData{
			MaxHP: 3,
			CurHP: 3,
			Team:  unit.TeamGood,
		},
	}
}

func (char *character) RecvDamage(dmg int, from unit.Unit) {
	unit.RecvDamageDefault(char, dmg, from)
}

func (char *character) Wearing() Item {
	return char.wearing
}

func (char *character) WearOrTakeOff(idx int) error {
	if idx >= len(char.Inventory()) {
		return fmt.Errorf("no such item")
	}
	item := char.Inventory()[idx]

	switch {
	case char.wearing == nil:
		err := item.Wear(char)
		if err != nil {
			return err
		}
		char.wearing = item
	case char.wearing == item:
		err := item.TakeOff(char)
		if err != nil {
			return err
		}
		char.wearing = nil
	default:
		return fmt.Errorf("you should take off other items first")
	}

	return nil
}

func (char *character) Inventory() []Item {
	return char.inventory
}

func (char *character) AddItem(item Item) {
	char.inventory = append(char.inventory, item)
}

func (char *character) Interact(other objects.GameObject) {
	unit.InteractDefault(char, other)
}

func (char *character) Response(other objects.GameObject) {
	// empty
}

func (char *character) ModelName() string {
	return "character"
}
