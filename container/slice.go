package container

import "sync"

type ConcurrentSlice[T any] struct {
	inner []T
	mutex sync.RWMutex
}

func NewConcurrentSlice[T any](capacity int) *ConcurrentSlice[T] {
	return &ConcurrentSlice[T]{
		inner: make([]T, 0, capacity),
		mutex: sync.RWMutex{},
	}
}

func (slice *ConcurrentSlice[T]) Append(elements ...T) {
	slice.mutex.Lock()
	defer slice.mutex.Unlock()
	slice.inner = append(slice.inner, elements...)
}

func (slice *ConcurrentSlice[T]) Len() int {
	slice.mutex.RLock()
	length := len(slice.inner)
	slice.mutex.RUnlock()
	return length
}

func (slice *ConcurrentSlice[T]) AsSlice() []T {
	slice.mutex.RLock()
	list := make([]T, 0, len(slice.inner))
	list = append(list, slice.inner...)
	//copy(list, slice.inner) // TODO: 貌似有点问题, https://stackoverflow.com/questions/30182538/why-cant-i-duplicate-a-slice-with-copy
	slice.mutex.RUnlock()

	return list
}
