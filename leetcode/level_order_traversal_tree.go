// problem #102 (https://leetcode.com/problems/binary-tree-level-order-traversal/)
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	var (
		result [][]int
	)

	for lvl := 0; ; lvl++ {
		a := make(acc, 0)
		a.process(root, 0, lvl)
		if len(a) == 0 {
			break
		}
		result = append(result, []int(a))
	}

	return result
}

type acc []int

func (a *acc) process(root *TreeNode, cl, hl int) {
	if root == nil {
		return
	}
	if cl == hl {
		*a = append(*a, root.Val)
		return
	}

	a.process(root.Left, cl+1, hl)
	a.process(root.Right, cl+1, hl)
}
