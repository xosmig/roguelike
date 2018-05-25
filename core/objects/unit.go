package objects

type Unit interface {
	GameObject
	HasPosition
	RecvDamage(dmg int, from Unit)
	GetTeam() int
	GetHP() int
	SetHP(hp int)
	GetMaxHP() int
	SetMaxHP(hp int)
	Die(from Unit)
}

type UnitData struct {
	PositionData
	MaxHP    int
	CurHP    int
	Team     int
}

func RecvDamageDefault(unit Unit, dmg int, from Unit) {
	newHP := unit.GetHP() - dmg
	unit.SetHP(newHP)
	if unit.GetHP() == 0 {
		unit.Die(from)
	}
}

func UnitInteractDefault(unit Unit, other GameObject) {
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
	return unit.CurHP
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

func (unit *UnitData) Blocks(other GameObject) bool {
	if otherUnit, ok := other.(Unit); ok {
		return unit.GetTeam() != otherUnit.GetTeam()
	}
	return false
}

func (unit *UnitData) IsAlive() bool {
	return unit.GetHP() > 0
}
