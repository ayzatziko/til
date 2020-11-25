package crackint

import (
	"fmt"
	"testing"
)

func TestWaysToRunUp(t *testing.T) {
	tt := []struct {
		stairsN int
		ways    int
	}{
		{}, {1, 1}, {2, 2}, {3, 4}, {4, 7}, {5, 13}, {6, 24},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d_stairs_num", tc.stairsN), func(t *testing.T) {
			t.Run("non_memo", func(t *testing.T) {
				ways := waysToRunUp(tc.stairsN)
				if tc.ways != ways {
					t.Fatalf("expected %d ways, got %d", tc.ways, ways)
				}
			})

			t.Run("memo", func(t *testing.T) {
				ways := waysToRunUpMemo(tc.stairsN)
				if tc.ways != ways {
					t.Fatalf("expected %d ways, got %d", tc.ways, ways)
				}
			})
		})
	}
}
