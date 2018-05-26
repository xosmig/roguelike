package items

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/character/mock_character"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/gamemap/mock_gamemap"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/state/mock_state"
	"testing"
)

func TestItemObject_Response(tt *testing.T) {
	ctrl := gomock.NewController(tt)

	char := mock_character.NewMockCharacter(ctrl)
	item := mock_character.NewMockItem(ctrl)
	gameMap := mock_gamemap.NewMockGameMap(ctrl)
	st := mock_state.NewMockGameState(ctrl)

	itemObj := NewItemObject(st, item)
	itemObj.SetPosition(geom.Loc(5, 5))

	st.EXPECT().GetMap().Return(gameMap)
	char.EXPECT().AddItem(item)
	cell := &gamemap.Cell{Object: itemObj}
	gameMap.EXPECT().Get(geom.Loc(5, 5)).Return(cell)

	itemObj.Response(char)
	assert.Equal(tt, objects.Empty, cell.Object)
}
