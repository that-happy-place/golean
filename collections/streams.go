package collections

func TransformSlice[F any, T any](transformer func(src F) T, sourceSlice []F) []T {
	if len(sourceSlice) == 0 {
		return nil
	}

	to := make([]T, 0, len(sourceSlice))
	for _, ele := range sourceSlice {
		to = append(to, transformer(ele))
	}

	return to
}

func FilterSlice[T any](filter func(src T) bool, sourceSlice []T) []T {
	if len(sourceSlice) == 0 {
		return nil
	}

	var to []T
	for _, ele := range sourceSlice {
		if filter(ele) {
			to = append(to, ele)
		}
	}

	return to
}
