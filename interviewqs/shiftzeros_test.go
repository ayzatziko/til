package interviewqs

import (
	"fmt"
	"testing"
)

func TestShiftzeros(t *testing.T) {
	tt := []struct {
		arr []int
		res []int
	}{
		{arr: []int{1, 2, 2, 0, 0, 7, 43, 0, 2, 0, 6}, res: []int{1, 2, 2, 7, 43, 2, 6, 0, 0, 0, 0}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.arr), func(t *testing.T) {
			orig := make([]int, len(tc.arr))
			copy(orig, tc.arr)
			t.Run("nonoptimized", func(t *testing.T) {
				shiftzeros(tc.arr)
				if len(tc.arr) != len(tc.res) {
					t.Fatalf("expected %v, got %v", tc.res, tc.arr)
				}

				for i, v := range tc.res {
					if v != tc.arr[i] {
						t.Fatalf("expected %v, got %v", tc.res, tc.arr)
					}
				}
			})

			tc.arr = orig
			t.Run("optimized", func(t *testing.T) {
				shiftzeros(tc.arr)
				if len(tc.arr) != len(tc.res) {
					t.Fatalf("expected %v, got %v", tc.res, tc.arr)
				}

				for i, v := range tc.res {
					if v != tc.arr[i] {
						t.Fatalf("expected %v, got %v", tc.res, tc.arr)
					}
				}
			})
		})
	}
}
