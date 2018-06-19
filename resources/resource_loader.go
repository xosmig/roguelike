package resources

import (
	"bytes"
	"fmt"
	"io"
)

// Loader is used to abstract loading game resources (from memory / disc / network / etc)
type Loader interface {
	Load(name string) (io.Reader, error)
}

// ErrNotFound should be returned by loaders when the resource with the given name cannot be found
var ErrNotFound = fmt.Errorf("resource not found")

// InMemoryLoader loads resources stored in memory
type InMemoryLoader struct {
	resources map[string][]byte
}

// NewInMemoryLoader creates an empty `InMemoryLoader`
func NewInMemoryLoader() *InMemoryLoader {
	return &InMemoryLoader{resources: make(map[string][]byte)}
}

// Load returns an io.Reader over the given string for resource `name`
func (loader *InMemoryLoader) Load(name string) (io.Reader, error) {
	data, present := loader.resources[name]
	if !present {
		return nil, ErrNotFound
	}
	return bytes.NewBuffer(data), nil
}

// AddString adds a new resource to `InMemoryLoader`
func (loader *InMemoryLoader) AddString(name string, data string) {
	loader.resources[name] = []byte(data)
}
