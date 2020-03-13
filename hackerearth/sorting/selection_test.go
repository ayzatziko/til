package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSortN(t *testing.T) {
	tests := []struct {
		in, out []int
		steps   int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 2},
		{[]int{2, 3, 1, 4, 5}, []int{1, 3, 2, 4, 5}, 1},
		{[]int{2, 3, 1, 4, 5}, []int{1, 2, 3, 4, 5}, 2},
	}

	for i, tt := range tests {
		o := SelectionSortN(tt.in, tt.steps)
		if !reflect.DeepEqual(tt.out, o) {
			t.Errorf("test %d: in %v, steps %d, want %v, got %v\n", i, tt.in, tt.steps, tt.out, o)
		}
	}
}
