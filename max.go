package mega

import "constraints"

func Max[T constraints.Ordered](sequence ...T) T {
	if len(sequence) == 0 {
		panic("arg is an empty sequence")
	}

	var max T
	for idx := 0; idx < len(sequence); idx++ {
		item := sequence[idx]
		if idx == 0 {
			max = item
			continue
		}
		if item > max {
			max = item
		}
	}
	return max
}
