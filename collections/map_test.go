package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Map(t *testing.T) {
	type args[K comparable, V any] struct {
		mapper func(k K, v V) (K, V)
	}
	type testCase[K comparable, V any] struct {
		name string
		m    Map[K, V]
		args args[K, V]
		want Map[K, V]
	}
	tests := []testCase[string, int]{
		{
			name: "normal case",
			m: map[string]int{
				"one":   0,
				"two":   1,
				"three": 2,
			},
			args: args[string, int]{
				mapper: func(k string, v int) (string, int) {
					return k, v + 1
				},
			},
			want: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			name: "nil case",
			m:    nil,
			args: args[string, int]{
				mapper: func(k string, v int) (string, int) {
					return k, v + 1
				},
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.m.Map(tt.args.mapper), "Map(%v)", tt.args.mapper)
		})
	}
}

func TestMap_ToKeySlice(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		m    Map[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{
			name: "normal case",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			want: []string{
				"one",
				"two",
				"three",
			},
		},
		{
			name: "nil case",
			m:    nil,
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, len(tt.want), len(tt.m.ToKeySlice()), "ToKeySlice()")
		})
	}
}

func TestMap_ToValueSlice(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		m    Map[K, V]
		want []V
	}
	tests := []testCase[string, int]{
		{
			name: "normal case",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			want: []int{
				1, 2, 3,
			},
		},
		{
			name: "nil case",
			m:    nil,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, len(tt.want), len(tt.m.ToValueSlice()), "ToValueSlice()")
		})
	}
}
