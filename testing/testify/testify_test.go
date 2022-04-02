package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}

func TestEmpty(t *testing.T) {
	assert.Empty(t, "", "should be empty")

	// assert for nil (good for errors)
	var empty []string = nil
	assert.Nil(t, empty)
	assert.NotNil(t, []string{})
}

func TestNotEmpty(t *testing.T) {
	assert.NotEmpty(t, "x", "should not be empty")

	var slice []string = []string{"a", "b"}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, slice) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "a", slice[0])
		assert.Equal(t, "b", slice[1])
	}
}

func TestError(t *testing.T) {

	// assert for error
	err := fmt.Errorf("error message")
	assert.Error(t, err)

	// assert for not error
	assert.NoError(t, nil)
}
