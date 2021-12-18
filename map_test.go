package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	actual := Map([]int{1, 2, 3, 4}, func(x int) string { return "Hello" })
	except := []string{"Hello", "Hello", "Hello", "Hello"}

	assert.Equal(t, except, actual)
}
