package generativerecursion

import (
	"reflect"
	"testing"
)

func TestCountingPyramid(t *testing.T) {
	tests := []struct{
		input int
		output [][]int
	}{
		{7, [][]int{[]int{1},[]int{1,1},[]int{2,1},[]int{1,2,1,1},[]int{1,1,1,2,2,1},[]int{3,1,2,2,1,1},[]int{1,3,1,1,2,2,2,1}}},
	}

	for i, tt := range tests {
		output := CountPyramid(tt.input)
		if !reflect.DeepEqual(tt.output, output) {
			t.Errorf("test %d: input %d, want %v, got %v\n", i, tt.input, tt.output, output)
		}
	}
}