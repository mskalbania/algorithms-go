package other

import (
	"cmp"
	"maps"
)

type Set[T cmp.Ordered] map[T]struct{}

func NewSet[T cmp.Ordered](vals ...T) Set[T] {
	out := make(map[T]struct{}, len(vals))
	for _, arg := range vals {
		out[arg] = struct{}{}
	}
	return out
}

func (s Set[T]) Add(vals ...T) {
	for _, val := range vals {
		s[val] = struct{}{}
	}
}

func (s Set[T]) Remove(entry T) bool {
	for k := range s {
		if k == entry {
			delete(s, k)
			return true
		}
	}
	return false
}

func (s Set[T]) Values() []T {
	seq := maps.Keys(s)
	out := make([]T, 0, len(s))
	for v := range seq {
		out = append(out, v)
	}
	return out
}
