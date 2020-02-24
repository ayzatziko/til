package leetcode

import "testing"

func TestSymmetricTree(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want bool
	}{
		// ok
		"base":     {nil, true},
		"not base": {&TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}, true},

		// fail
		"not symmetric": {&TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, false},
	}

	for name, tt := range tests {
		got := IsSymmetric(tt.root)
		if tt.want != got {
			t.Errorf("test %s: want %v got %v", name, tt.want, got)
		}
	}
}
