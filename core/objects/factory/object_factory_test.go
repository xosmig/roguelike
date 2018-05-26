package factory

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/xosmig/roguelike/core/objects/mock_objects"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	t.Run("ReturnsExpectedObject", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		obj := mock_objects.NewMockGameObject(ctrl)
		f := Singleton(obj)

		created, err := f.Create(geom.Loc(5, 5))
		assert.NoError(tt, err)
		assert.Equal(tt, obj, created)
	})

	t.Run("SetsPosition", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		obj := mock_objects.NewMockMovableObject(ctrl)
		f := Singleton(obj)

		obj.EXPECT().SetPosition(geom.Loc(5, 5))
		_, err := f.Create(geom.Loc(5, 5))
		assert.NoError(tt, err)
	})

	t.Run("DoesntReturnObjectTwice", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		obj := mock_objects.NewMockGameObject(ctrl)
		f := Singleton(obj)

		_, err := f.Create(geom.Loc(5, 5))
		assert.NoError(tt, err)

		_, err = f.Create(geom.Loc(2, 2))
		assert.Error(tt, err)
	})
}

func TestRepeated(t *testing.T) {
	t.Run("ReturnsSameObject", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		obj := mock_objects.NewMockGameObject(ctrl)
		f := Repeated(obj)

		created1, err := f.Create(geom.Loc(5, 5))
		assert.NoError(tt, err)
		assert.Equal(tt, obj, created1)

		created2, err := f.Create(geom.Loc(5, 5))
		assert.NoError(tt, err)
		assert.Equal(tt, obj, created2)
	})
}
