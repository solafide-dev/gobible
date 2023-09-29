package gobible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	g := NewGoBible()
	err := g.Load("data/WEB.json")
	assert.Nil(t, err)
}

func TestLoadFormat(t *testing.T) {
	g := NewGoBible()
	err := g.LoadFormat("data/WEB.xml", "osis")
	assert.Nil(t, err)
}

func TestLoadString(t *testing.T) {
	g := NewGoBible()
	err := g.LoadString(`{"version": {"name": "Legacy Standard Bible", "abbrev": "LSB"}, "books": []}`)
	assert.Nil(t, err)
}

func TestGetBibleJSON(t *testing.T) {
	g := NewGoBible()
	err := g.Load("data/WEB.json")
	assert.Nil(t, err)

	// Convert WEB to a string
	bJSON, err := g.GetBibleJSON("WEB")
	assert.Nil(t, err)
	assert.NotEmpty(t, bJSON)

	// Load the string back into gobible
	// This should generate an "already loaded" error
	err = g.LoadString(bJSON)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "already loaded")
	}

	// Unload WEB so we can try again
	err = g.Unload("WEB")
	assert.Nil(t, err)

	// Unload WEB against to cover the error case
	err = g.Unload("WEB")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "not found")
	}

	// Load the string back into gobible
	err = g.LoadString(bJSON)
	assert.Nil(t, err)
}
