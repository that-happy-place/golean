package collections

type Collection[T comparable] interface {
	Add(T) bool
	AddAll(Collection[T]) bool
	Clear()
	Contains(T) bool
	ContainsAll(Collection[T]) bool
	IsEmpty() bool
	Remove(T) bool
	RemoveAll(Collection[T]) bool
	RemoveIf(func(T) bool) bool
	Len() int
	ToSlice() []T
}

func NewCollectionFromSlice[T comparable](s []T) Collection[T] {
	return &sliceCollection[T]{s: s}
}

type sliceCollection[T comparable] struct {
	s []T
}

func (s *sliceCollection[T]) Add(val T) bool {
	s.s = append(s.s, val)
	return true
}

func (s *sliceCollection[T]) AddAll(vals Collection[T]) bool {
	s.s = append(s.s, vals.ToSlice()...)
	return true
}

func (s *sliceCollection[T]) Clear() {
	s.s = nil
}

func (s *sliceCollection[T]) Contains(val T) bool {
	for _, v := range s.s {
		if v == val {
			return true
		}
	}

	return false
}

func (s *sliceCollection[T]) ContainsAll(vals Collection[T]) bool {
	containsAll := true
	for _, val := range vals.ToSlice() {
		containsAll = containsAll && s.Contains(val)
	}
	return containsAll
}

func (s *sliceCollection[T]) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *sliceCollection[T]) Remove(val T) bool {
	index := -1
	for i, v := range s.s {
		if v == val {
			index = i
			break
		}
	}

	if index >= 0 {
		s.s = append(s.s[:index], s.s[index+1:]...)
	}

	return index != -1
}

func (s *sliceCollection[T]) RemoveAll(vals Collection[T]) bool {
	changed := false
	for _, val := range vals.ToSlice() {
		changed = s.Remove(val) || changed
	}
	return changed
}

func (s *sliceCollection[T]) RemoveIf(predicate func(T) bool) bool {
	changed := false
	for _, val := range s.ToSlice() {
		if predicate(val) {
			changed = s.Remove(val) || changed
		}
	}
	return changed
}

func (s *sliceCollection[T]) Len() int {
	return len(s.s)
}

func (s *sliceCollection[T]) ToSlice() []T {
	copied := make([]T, len(s.s))
	copy(copied, s.s)
	return copied
}
