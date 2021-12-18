package mega

import "sync"

// Map manipulates an iterator and transforms it to another type.
func Map[T1 any, T2 any](sequence []T1, f func(x T1) T2) []T2 {
	result := make([]T2, 0, len(sequence))
	for _, x := range sequence {
		result = append(result, f(x))
	}
	return result
}

func ParallelMap[T1 any, T2 any](sequence []T1, f func(x T1) T2) []T2 {
	ch := make(chan T2)

	go func() {
		defer close(ch)

		var wg sync.WaitGroup
		wg.Add(len(sequence))
		for _, x := range sequence {
			x := x
			go func() {
				defer wg.Done()
				ch <- f(x)
			}()
		}
		wg.Wait()
	}()

	result := make([]T2, 0, len(sequence))
	for x := range ch {
		result = append(result, x)
	}
	return result
}
