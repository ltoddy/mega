package mega

func Chunk[T any](sequence []T, size int) [][]T {
	var chunks [][]T
	for size < len(sequence) {
		sequence, chunks = sequence[size:], append(chunks, sequence[0:size:size])
	}
	return append(chunks, sequence)
}
