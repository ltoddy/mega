package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipEmptyResult(t *testing.T) {
	array1 := []int{21, 22, 23}
	emptySlice := []int{}

	expected := []Tuple[int, int]{}
	result := Zip(emptySlice, array1)
	assert.Equal(t, result, expected)
}

func TestZipSlices(t *testing.T) {
	slice1 := []int{11, 12, 13}
	slice2 := []int{21, 22, 23, 24, 25}

	expected := []Tuple[int, int]{
		{First: 11, Second: 21},
		{First: 12, Second: 22},
		{First: 13, Second: 23},
	}

	result := Zip(slice1, slice2)
	assert.Equal(t, result, expected)
}

func TestZipArrays(t *testing.T) {
	array1 := [...]int{11, 12, 13}
	array2 := [...]int{21, 22, 23, 24, 25}

	expected := []Tuple[int, int]{
		{First: 11, Second: 21},
		{First: 12, Second: 22},
		{First: 13, Second: 23},
	}
	result := Zip(array1[:], array2[:])
	assert.Equal(t, result, expected)
}
