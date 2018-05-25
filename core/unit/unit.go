package unit

import (
	"github.com/xosmig/roguelike/core/objects"
	"log"
)

type Unit interface {
	objects.MovableObject
	RecvDamage(dmg int, from Unit)
	GetTeam() int
	GetHP() int
	SetHP(hp int)
	GetMaxHP() int
	SetMaxHP(hp int)
}

const (
	TeamGood = 1
	TeamEvil = 2
)

type UnitData struct {
	objects.PositionData
	MaxHP int
	CurHP int
	Team  int
}

func RecvDamageDefault(unit Unit, dmg int, from Unit) {
	newHP := unit.GetHP() - dmg
	unit.SetHP(newHP)
	log.Printf("Debug: %T damaged by %T, hp=%d\n", unit, from, unit.GetHP())
}

func InteractDefault(unit Unit, other objects.GameObject) {
	if otherUnit, ok := other.(Unit); ok {
		if otherUnit.GetTeam() != unit.GetTeam() {
			otherUnit.RecvDamage(1, unit)
		}
	}
}

func (unit *UnitData) GetTeam() int {
	return unit.Team
}

func (unit *UnitData) GetHP() int {
	return unit.CurHP
}

func (unit *UnitData) checkHPBounds() {
	if unit.CurHP < 0 {
		unit.CurHP = 0
	}
	if unit.CurHP > unit.MaxHP {
		unit.CurHP = unit.MaxHP
	}
}

func (unit *UnitData) SetHP(hp int) {
	unit.CurHP = hp
	unit.checkHPBounds()
}

func (unit *UnitData) GetMaxHP() int {
	return unit.MaxHP
}

func (unit *UnitData) SetMaxHP(hp int) {
	unit.MaxHP = hp
	unit.checkHPBounds()
}

func (unit *UnitData) SetMaxHPPreserveNorm(hp int) {
	hpNorm := unit.GetHPNorm()
	unit.MaxHP = hp
	unit.SetHPNorm(hpNorm)
}

func (unit *UnitData) GetHPNorm() float32 {
	return float32(unit.GetHP()) / float32(unit.GetMaxHP())
}

func (unit *UnitData) SetHPNorm(hpNorm float32) {
	unit.SetHP(int(float32(unit.GetMaxHP())*hpNorm + 0.5))
}

func IsAlive(unit Unit) bool {
	return unit.GetHP() > 0
}
