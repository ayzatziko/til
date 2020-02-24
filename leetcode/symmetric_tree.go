// problem #101 (https://leetcode.com/problems/symmetric-tree/)
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return isSymm(root, root)
}

func isSymm(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil {
		if left.Val == right.Val {
			return isSymm(left.Left, right.Right) && isSymm(left.Right, right.Left)
		}

		return false
	}

	return false
}
