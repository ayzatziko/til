package leetcode

import (
	"reflect"
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	tests := []struct {
		input int
		want  [][]int
	}{
		{0, nil},
		{1, [][]int{[]int{1}}},
		{5, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}},
	}

	for _, tt := range tests {
		if !reflect.DeepEqual(PascalsTriangle(tt.input), tt.want) {
			t.Errorf("invalid output with input %d\n", tt.input)
		}
	}
}
