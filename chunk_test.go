package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	sequence := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	chunks := Chunk(sequence, 4)

	assert.Equal(t, 3, len(chunks))
	assert.Equal(t, []int64{1, 2, 3, 4}, chunks[0])
	assert.Equal(t, []int64{5, 6, 7, 8}, chunks[1])
	assert.Equal(t, []int64{9}, chunks[2])
}
