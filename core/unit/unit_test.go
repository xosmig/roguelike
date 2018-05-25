package unit

import (
	"testing"
	"github.com/xosmig/roguelike/core/objects"
)

func TestUnitDataIsNotGameObject(t *testing.T) {
	var raw interface{} = &UnitData{}
	if _, ok := raw.(objects.GameObject); ok {
		t.Errorf("UnitData should not implement Unit by itself")
	}
}
