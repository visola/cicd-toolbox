package arrayutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	arr := []string{"one", "two", "three", "four", "five"}

	assert.Contains(t, arr, "two")
	assert.Contains(t, arr, "four")

	assert.NotContains(t, arr, "six")
	assert.NotContains(t, arr, "zero")
}
