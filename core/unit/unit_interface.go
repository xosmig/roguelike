package unit

import (
	"github.com/xosmig/roguelike/core/objects"
	"log"
)

// Unit describes battle unit. Including character and monster
type Unit interface {
	objects.MovableObject
	RecvDamage(dmg int, from Unit)
	GetTeam() int
	GetHP() int
	SetHP(hp int)
	GetMaxHP() int
	SetMaxHP(hp int)
}

// Might be more teams in future
const (
	TeamGood = 1
	TeamEvil = 2
)

// RecvDamageDefault is the default implementation of `RecvDamage` method.
func RecvDamageDefault(unit Unit, dmg int, from Unit) {
	unit.SetHP(unit.GetHP() - dmg)
	log.Printf("Debug: %T damaged by %T, hp=%d\n", unit, from, unit.GetHP())
}

// InteractDefault is the default implementation of `Interact` method for battle units.
// Will attack the other object if it is an enemy unit.
func InteractDefault(unit Unit, other objects.GameObject) {
	if otherUnit, ok := other.(Unit); ok {
		if otherUnit.GetTeam() != unit.GetTeam() {
			otherUnit.RecvDamage(1, unit)
		}
	}
}

// IsAlive returns true if the unit has more than 0 health points
func IsAlive(unit Unit) bool {
	return unit.GetHP() > 0
}
