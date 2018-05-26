// Tests have to be placed in a separate package to avoid dependency cycle between `character` and `mock_character`
package tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/character/mock_character"
	"testing"
)

func TestInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item0 := mock_character.NewMockItem(ctrl)
	item1 := mock_character.NewMockItem(ctrl)

	assert.Empty(t, char.Inventory())
	char.AddItem(item0)
	char.AddItem(item1)

	assert.Len(t, char.Inventory(), 2)
	assert.Contains(t, char.Inventory(), item0)
	assert.Contains(t, char.Inventory(), item1)
}

func TestCantWear2Items(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item0 := mock_character.NewMockItem(ctrl)
	item1 := mock_character.NewMockItem(ctrl)
	char.AddItem(item0)
	char.AddItem(item1)

	item0.EXPECT().Wear(char).Times(1).Return(nil)

	assert.NoError(t, char.WearOrTakeOff(0))
	assert.Error(t, char.WearOrTakeOff(1))
}

func TestSuccessfulWearAndTakeOff(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item := mock_character.NewMockItem(ctrl)
	char.AddItem(item)

	item.EXPECT().Wear(char).Times(1).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(0))
	assert.Equal(t, item, char.Wearing())

	item.EXPECT().TakeOff(char).Times(1).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(0))
	assert.Equal(t, nil, char.Wearing())
}

func TestFailedWearing(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item := mock_character.NewMockItem(ctrl)
	char.AddItem(item)

	item.EXPECT().Wear(char).Times(1).Return(errors.New("mock"))
	assert.Error(t, char.WearOrTakeOff(0))
	assert.Equal(t, nil, char.Wearing())
}

func TestFailedTakeOff(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item := mock_character.NewMockItem(ctrl)
	char.AddItem(item)

	item.EXPECT().Wear(char).Times(1).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(0))
	assert.Equal(t, item, char.Wearing())

	item.EXPECT().TakeOff(char).Times(1).Return(errors.New("mock"))
	assert.Error(t, char.WearOrTakeOff(0))
	assert.Equal(t, item, char.Wearing())
}

func TestWearAnotherItemAfterTakeOff(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item0 := mock_character.NewMockItem(ctrl)
	item1 := mock_character.NewMockItem(ctrl)
	char.AddItem(item0)
	char.AddItem(item1)

	item0.EXPECT().Wear(char).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(0))
	item0.EXPECT().TakeOff(char).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(0))
	item1.EXPECT().Wear(char).Return(nil)
	assert.NoError(t, char.WearOrTakeOff(1))
	assert.Equal(t, item1, char.Wearing())
}
