package math

import (
	"math"
)

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

func roundToFloor(num float64) int {
	return int(num)
}

func roundToCeil(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func FloorToPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(roundToFloor(num * output)) / output
}

func CeilToPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(roundToCeil(num * output)) / output
}
