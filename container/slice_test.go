package container

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentSlice(t *testing.T) {
	slice := NewConcurrentSlice(16)

	for i := 0; i < 100; i++ {
		go slice.Append(1, 2, 3)
	}
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 300, slice.Len())
}
