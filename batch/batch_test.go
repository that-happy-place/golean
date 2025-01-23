package batch

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args[T any] struct {
		array     []T
		batchSize uint32
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "Test case 1",
			args: args[string]{
				array:     []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
				batchSize: 3,
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
				{"j"},
			},
		},
		{
			name: "Test case 2",
			args: args[string]{
				array:     []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"},
				batchSize: 3,
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
				{"j", "k", "l"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.array, tt.args.batchSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
