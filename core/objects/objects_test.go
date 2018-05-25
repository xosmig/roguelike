package objects

import (
	"testing"
)

func TestUnitDataIsNotGameObject(t *testing.T) {
	var raw interface{} = &UnitData{}
	if _, ok := raw.(GameObject); ok {
		t.Errorf("UnitData should not implement Unit by itself")
	}
}
