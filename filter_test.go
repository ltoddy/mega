package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	r := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })

	assert.Equal(t, []int{2, 4}, r)
}

func TestParallelFilter(t *testing.T) {
	r := ParallelFilter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })

	assert.Equal(t, 2, len(r))
}
