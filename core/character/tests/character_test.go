package tests

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/character/mock_character"
	"github.com/xosmig/roguelike/core/character"
	"errors"
)

func TestCharacter_Inventory(t *testing.T) {
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

func TestCharacter_CantWear2Items(t *testing.T) {
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

func TestCharacter_SuccessfulWearAndTakeOff(t *testing.T) {
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

func TestCharacter_FailedWearing(t *testing.T) {
	ctrl := gomock.NewController(t)
	char := character.New()

	item := mock_character.NewMockItem(ctrl)
	char.AddItem(item)

	item.EXPECT().Wear(char).Times(1).Return(errors.New("mock"))
	assert.Error(t, char.WearOrTakeOff(0))
	assert.Equal(t, nil, char.Wearing())
}

func TestCharacter_FailedTakeOff(t *testing.T) {
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

func TestCharacter_WearAnotherItemAfterTakeOff(t *testing.T) {
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
