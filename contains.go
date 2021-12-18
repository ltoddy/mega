package mega

// Contains returns true if an element is present in a iterator.
func Contains[T comparable](sequence []T, element T) bool {
	for _, x := range sequence {
		if x == element {
			return true
		}
	}
	return false
}
