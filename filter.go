package mega

import "sync"

func Filter[T any](sequence []T, predicate func(x T) bool) []T {
	result := make([]T, 0, len(sequence))
	for _, x := range sequence {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

func ParallelFilter[T any](sequence []T, predicate func(x T) bool) []T {
	ch := make(chan T)

	go func() {
		defer close(ch)

		var wg sync.WaitGroup
		wg.Add(len(sequence))
		for _, x := range sequence {
			x := x
			go func() {
				defer wg.Done()
				if predicate(x) {
					ch <- x
				}
			}()
		}
		wg.Wait()
	}()

	result := make([]T, 0, len(sequence))
	for x := range ch {
		result = append(result, x)
	}
	return result
}
