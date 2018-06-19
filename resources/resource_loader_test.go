package resources

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestInMemoryLoader_Load(t *testing.T) {
	data := "hello\nworld"
	loader := NewInMemoryLoader()

	loader.AddString("foo/bar", data)

	reader, err := loader.Load("foo/bar")
	assert.NoError(t, err)
	bytes, err := ioutil.ReadAll(reader)
	assert.NoError(t, err)
	assert.Equal(t, data, string(bytes))

	_, err = loader.Load("bar/baz")
	assert.Error(t, err)
}
