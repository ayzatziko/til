package abstractions

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		in, o []int
		f     func(int, int) bool
	}{
		{nil, nil, nil},
		{[]int{1, 2, 3}, []int{3, 2, 1}, func(a, b int) bool { return a < b }},
		{[]int{3, 2, 1}, []int{1, 2, 3}, func(a, b int) bool { return a > b }},
	}

	for i, tt := range tests {
		o := Sort(tt.in, tt.f)
		if !reflect.DeepEqual(tt.o, o) {
			t.Errorf("test %d: want %v, got %v\n", i, tt.o, o)
		}
	}
}
