package batch

func Split[T any](array []T, batchSize uint32) [][]T {
	if batchSize == 0 {
		panic("batchSize cannot be 0")
	}
	var batches [][]T
	for batchSize < uint32(len(array)) {
		array, batches = array[batchSize:], append(batches, array[0:batchSize:batchSize])
	}
	return append(batches, array)
}
