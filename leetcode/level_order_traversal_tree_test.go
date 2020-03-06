// problem #128 (https://leetcode.com/problems/longest-consecutive-sequence/)
package leetcode

import "testing"

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		// ok
		{nil, nil},
		{
			&TreeNode{
				Val: 1, Left: &TreeNode{
					Val: 3, Left: &TreeNode{Val: 4},
				}, Right: &TreeNode{
					Val: 3, Right: &TreeNode{Val: 4}},
			}, [][]int{[]int{1}, []int{3, 3}, []int{4, 4}},
		},
	}

	for i, tt := range tests {
		got := LevelOrder(tt.root)
		if len(got) != len(tt.want) {
			t.Errorf("test %d: got len %d, want len %d\n", i, len(got), len(tt.want))
		}
		for lvl := range got {
			if len(got[lvl]) != len(tt.want[lvl]) {
				t.Errorf("test %d: level %d, got len %d, want len %d\n", i, lvl, len(got[lvl]), len(tt.want[lvl]))
			}
			for idx := range got[lvl] {
				vgot := got[lvl][idx]
				vwant := tt.want[lvl][idx]
				if vgot != vwant {
					t.Errorf("test %d: level %d, value index %d: want %d, got %d\n", i, lvl, idx, vgot, vwant)
				}
			}
		}
	}
}
