package mega

// Compact creates a slice with all empty/zero values removed.
func Compact[T comparable](sequence []T) []T {
	var zero T

	result := make([]T, 0, len(sequence))
	for _, elem := range sequence {
		if elem == zero {
			continue
		}
		result = append(result, elem)
	}
	return result
}
