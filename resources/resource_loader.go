package resources

import (
	"io"
	"fmt"
	"bytes"
)

type Loader interface {
	Load(name string) (io.Reader, error)
}

type InMemoryLoader struct {
	resources map[string][]byte
}

var NotFound = fmt.Errorf("resource not found")

func NewInMemoryLoader() *InMemoryLoader {
	return &InMemoryLoader{resources: make(map[string][]byte)}
}

func (loader *InMemoryLoader) Load(name string) (io.Reader, error) {
	data, present := loader.resources[name]
	if !present {
		return nil, NotFound
	}
	return bytes.NewBuffer(data), nil
}

func (loader *InMemoryLoader) AddString(name string, data string) {
	loader.resources[name] = []byte(data)
}
