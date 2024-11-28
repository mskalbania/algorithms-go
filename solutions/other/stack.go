package other

import (
	"errors"
	"slices"
	"sync"
)

var ErrStackEmpty = errors.New("stack empty")

var maxUnusedSpace = 50

type Stack[T any] struct {
	arr  []T
	lock sync.Mutex
}

func (s *Stack[T]) Push(val T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, val)
}

func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.arr) < 1 {
		var empty T
		return empty, ErrStackEmpty
	}
	last := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	if cap(s.arr)-len(s.arr) > maxUnusedSpace {
		s.arr = slices.Clip(s.arr)
	}
	return last, nil
}
