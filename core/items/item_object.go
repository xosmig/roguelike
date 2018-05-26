package items

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/state"
	"github.com/xosmig/roguelike/core/gamemap"
)

type ItemObject struct {
	objects.PositionData
	item character.Item
	st   state.GameState
}

func NewItemObject(st state.GameState, item character.Item) *ItemObject {
	return &ItemObject{
		st: st,
		item: item,
	}
}

func (obj *ItemObject) Response(other objects.GameObject) {
	if char, ok := other.(character.Character); ok {
		char.AddItem(obj.item)
		gamemap.Remove(obj.st.GetMap(), obj.GetPosition())
	}
}

func (obj *ItemObject) ModelName() string {
	return "item"
}
