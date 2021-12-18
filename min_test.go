package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWithArrayNumericInput(t *testing.T) {
	d1 := []int{8, 3, 4, 44, 0}
	n1 := []int{}
	d2 := []int8{3, 3, 5, 9, 1}
	n2 := []int8{}
	d3 := []int16{4, 5, 4, 33, 2}
	n3 := []int16{}
	d4 := []int32{5, 3, 21, 15, 3}
	n4 := []int32{}
	d5 := []int64{9, 3, 9, 1, 2}
	n5 := []int64{}

	r1 := Min(d1...)
	r2 := Min(d2...)
	r3 := Min(d3...)
	r4 := Min(d4...)
	r5 := Min(d5...)

	assert.Equal(t, int(0), r1, "It should return the min value in array")
	assert.Panics(t, func() { Min(n1...) }, "It should panic")
	assert.Equal(t, int8(1), r2, "It should return the min value in array")
	assert.Panics(t, func() { Min(n2...) }, "It should panic")
	assert.Equal(t, int16(2), r3, "It should return the min value in array")
	assert.Panics(t, func() { Min(n3...) }, "It should panic")
	assert.Equal(t, int32(3), r4, "It should return the min value in array")
	assert.Panics(t, func() { Min(n4...) }, "It should panic")
	assert.Equal(t, int64(1), r5, "It should return the min value in array")
	assert.Panics(t, func() { Min(n5...) }, "It should panic")

}

func TestMinWithArrayFloatInput(t *testing.T) {
	d1 := []float64{2, 38.3, 4, 4.4, 4}
	n1 := []float64{}
	d2 := []float32{2.9, 1.3, 4.23, 4.4, 7.7}
	n2 := []float32{}

	r1 := Min(d1...)
	r2 := Min(d2...)

	assert.Equal(t, float64(2), r1, "It should return the min value in array")
	assert.Panics(t, func() { Min(n1...) }, "It should panic")
	assert.Equal(t, float32(1.3), r2, "It should return the min value in array")
	assert.Panics(t, func() { Min(n2...) }, "It should panic")
}

func TestMinWithArrayInputWithStrings(t *testing.T) {
	d1 := []string{"abc", "abd", "cbd"}
	d2 := []string{"abc", "abd", "abe"}
	d3 := []string{"abc", "foo", " "}
	d4 := []string{"abc", "abc", "aaa"}
	n1 := []string{}

	r1 := Min(d1...)
	r2 := Min(d2...)
	r3 := Min(d3...)
	r4 := Min(d4...)

	assert.Equal(t, "abc", r1, "It should print cbd because its first char is min in the list")
	assert.Equal(t, "abc", r2, "It should print abe because its first different char is min in the list")
	assert.Equal(t, " ", r3, "It should print foo because its first different char is min in the list")
	assert.Equal(t, "aaa", r4, "It should print abc because its first different char is min in the list")
	assert.Panics(t, func() { Min(n1...) }, "It should panic")
}