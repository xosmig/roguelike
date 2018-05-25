package factory

import (
	"fmt"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/geom"
)

type ObjectFactory interface {
	Create(geom.Location) (objects.GameObject, error)
}

type repeatedObjectFactory struct {
	obj objects.GameObject
}

func Repeated(obj objects.GameObject) ObjectFactory {
	return repeatedObjectFactory{obj}
}

func (f repeatedObjectFactory) Create(pos geom.Location) (objects.GameObject, error) {
	// since it repeats the same objects in many places, it doesn't make sense to call SetPosition
	return f.obj, nil
}

type singletonObjectFactory struct {
	obj objects.GameObject
}

func Singleton(obj objects.GameObject) ObjectFactory {
	f := singletonObjectFactory{obj: obj}
	return f
}

func (f singletonObjectFactory) Create(pos geom.Location) (objects.GameObject, error) {
	if f.obj == nil {
		return nil, fmt.Errorf("double access to singleton object factory")
	}

	res := f.obj
	f.obj = nil

	if obj, ok := res.(objects.HasPosition); ok {
		obj.SetPosition(pos)
	}
	return res, nil
}
