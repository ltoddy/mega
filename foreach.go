package mega

// ForEach iterates over elements of collection and invokes iterator
// for each element.
func ForEach[T any](sequence []T, f func(x T)) {
	for _, x := range sequence {
		f(x)
	}
}
