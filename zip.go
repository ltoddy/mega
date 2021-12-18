package mega

type Tuple[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func Zip[T1 any, T2 any](sequence1 []T1, sequence2 []T2) []Tuple[T1, T2] {
	minLength := Min(len(sequence1), len(sequence2))

	result := make([]Tuple[T1, T2], 0, minLength)
	for i := 0; i < minLength; i++ {
		newTuple := Tuple[T1, T2]{
			First:  sequence1[i],
			Second: sequence2[i],
		}
		result = append(result, newTuple)
	}

	return result
}
