package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
	"testing"
)

func TestUnitDataIsNotGameObject(t *testing.T) {
	var raw interface{} = &unit.UnitData{}
	if _, ok := raw.(objects.GameObject); ok {
		t.Errorf("UnitData should not implement Unit by itself")
	}
}

func TestCheckHPBounds(t *testing.T) {
	ud := unit.UnitData{MaxHP: 3, CurHP: 2}
	ud.SetHP(-5)
	assert.Equal(t, 0, ud.GetHP())

	ud.SetHP(5)
	assert.Equal(t, 3, ud.GetHP())

	ud.SetMaxHP(2)
	assert.Equal(t, 2, ud.GetHP())
}
