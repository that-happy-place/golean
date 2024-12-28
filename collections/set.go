package collections

type Set[T comparable] interface {
	Collection[T]
	RetainAll(Collection[T]) bool
}

type HashSet[T comparable] struct {
	m map[T]nothing
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{m: make(map[T]nothing)}
}

func (h *HashSet[T]) Add(val T) bool {
	if _, ok := h.m[val]; ok {
		return false
	}
	h.m[val] = nothing{}
	return true
}

func (h *HashSet[T]) AddAll(vals Collection[T]) bool {
	changed := false
	for _, val := range vals.ToSlice() {
		changed = h.Add(val) || changed
	}
	return changed
}

func (h *HashSet[T]) Clear() {
	h.m = make(map[T]nothing)
}

func (h *HashSet[T]) Contains(val T) bool {
	_, ok := h.m[val]
	return ok
}

func (h *HashSet[T]) ContainsAll(vals Collection[T]) bool {
	containsAll := false
	for _, val := range vals.ToSlice() {
		containsAll = containsAll && h.Contains(val)
	}
	return containsAll
}

func (h *HashSet[T]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *HashSet[T]) Remove(val T) bool {
	_, ok := h.m[val]
	delete(h.m, val)
	return ok
}

func (h *HashSet[T]) RemoveAll(vals Collection[T]) bool {
	changed := false
	for _, val := range vals.ToSlice() {
		changed = h.Remove(val) || changed
	}
	return changed
}

func (h *HashSet[T]) RemoveIf(predicate func(T) bool) bool {
	changed := false
	for _, val := range h.ToSlice() {
		if predicate(val) {
			changed = h.Remove(val) || changed
		}
	}
	return changed
}

func (h *HashSet[T]) RetainAll(vals Collection[T]) bool {
	changed := false
	for _, val := range h.ToSlice() {
		if !vals.Contains(val) {
			changed = h.Remove(val) || changed
		}
	}
	return changed
}

func (h *HashSet[T]) Len() int {
	return len(h.m)
}

func (h *HashSet[T]) ToSlice() []T {
	resp := make([]T, len(h.m))
	index := 0
	for val := range h.m {
		resp[index] = val
		index++
	}
	return resp
}
