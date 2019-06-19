package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type S struct {
	Str string
	Int int
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var x interface{} = nil
	assert.Nil(t, x)

	object := S{Str: "Something"}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Str)

	}

}
