package collections

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Map(mapper func(k K, v V) (K, V)) Map[K, V] {
	newMap := make(map[K]V)
	if len(m) == 0 {
		return newMap
	}
	for k, v := range m {
		newK, newV := mapper(k, v)
		newMap[newK] = newV
	}
	return newMap
}

func (m Map[K, V]) ToMap() map[K]V {
	return m
}

func (m Map[K, V]) ToKeySlice() []K {
	keys := make([]K, 0, len(m))
	if len(m) == 0 {
		return keys
	}

	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) ToValueSlice() []V {
	values := make([]V, 0, len(m))
	if len(m) == 0 {
		return values
	}

	for _, v := range m {
		values = append(values, v)
	}
	return values
}
