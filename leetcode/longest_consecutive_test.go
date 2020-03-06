package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
	}

	for _, tt := range tests {
		got := LongestConsecutive(tt.input)
		if tt.want != got {
			t.Fatalf("want %d, got %d on %v\n", tt.want, got, tt.input)
		}
	}
}
