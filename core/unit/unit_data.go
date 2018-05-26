package unit

import "github.com/xosmig/roguelike/core/objects"

// UnitData contains properties, shared by most units. Not a game object.
type UnitData struct {
	objects.PositionData
	MaxHP int
	CurHP int
	Team  int
}

func (unit *UnitData) checkHPBounds() {
	if unit.CurHP < 0 {
		unit.CurHP = 0
	}
	if unit.CurHP > unit.MaxHP {
		unit.CurHP = unit.MaxHP
	}
}

// GetTeam returns the team of the unit
func (unit *UnitData) GetTeam() int {
	return unit.Team
}

// GetHP returns the current amount of health points of the unit
func (unit *UnitData) GetHP() int {
	return unit.CurHP
}

// SetHP sets the current amount of health points of the unit
func (unit *UnitData) SetHP(hp int) {
	unit.CurHP = hp
	unit.checkHPBounds()
}

// GetMaxHP returns the maximum amount of health points for the unit
func (unit *UnitData) GetMaxHP() int {
	return unit.MaxHP
}

// SetMaxHP sets the maximum amount of health points of the unit
func (unit *UnitData) SetMaxHP(hp int) {
	unit.MaxHP = hp
	unit.checkHPBounds()
}

// IsAlive returns true if the unit has more than 0 health points
func IsAlive(unit Unit) bool {
	return unit.GetHP() > 0
}
