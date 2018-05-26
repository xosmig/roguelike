package factory

import (
	"fmt"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
)

// ObjectFactory is primarily used by map loader to create objects.
type ObjectFactory interface {
	Create(geom.Location) (objects.GameObject, error)
}

type repeatedObjectFactory struct {
	obj objects.GameObject
}

// Repeated creates an object factory that can be called multiple times.
// `Create` never returns an error.
// Since it repeats the same objects in many places, it doesn't make sense to call `SetPosition`.
func Repeated(obj objects.GameObject) ObjectFactory {
	return repeatedObjectFactory{obj}
}

func (f repeatedObjectFactory) Create(pos geom.Location) (objects.GameObject, error) {
	return f.obj, nil
}

type singletonObjectFactory struct {
	obj objects.GameObject
}

// Singleton creates an object factory that can be called only 1 time.
// Further calls to `Create` will return errors.
// Calls `obj.SetPosition` if the object implements `HasPosition`
func Singleton(obj objects.GameObject) ObjectFactory {
	return &singletonObjectFactory{obj: obj}
}

func (f *singletonObjectFactory) Create(pos geom.Location) (objects.GameObject, error) {
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
