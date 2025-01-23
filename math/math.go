package math

func Min[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
