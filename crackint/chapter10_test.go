package crackint

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tt := []struct {
		in, out []int
	}{
		{nil, nil}, {[]int{1, 2, 3}, []int{1, 2, 3}}, {[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v=>%v", tc.in, tc.out), func(t *testing.T) {
			mergeSort(tc.in)
			if len(tc.out) != len(tc.in) {
				t.Fatalf("want %v, got %v", tc.out, tc.in)
			}

			for i, v := range tc.out {
				if v != tc.in[i] {
					t.Fatalf("want %v, got %v", tc.out, tc.in)
				}
			}
		})
	}
}
