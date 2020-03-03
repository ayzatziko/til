// problem #108 (https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
package leetcode

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	root := TreeNode{Val: nums[m]}
	insertLeft(&root, nums[:m])
	insertRight(&root, nums[m+1:])
	return &root
}

func insertLeft(root *TreeNode, nums []int) {
	if len(nums) == 0 {
		return
	}

	m := len(nums) / 2
	n := TreeNode{Val: nums[m]}
	insertLeft(&n, nums[:m])
	insertRight(&n, nums[m+1:])
	root.Left = &n
}

func insertRight(root *TreeNode, nums []int) {
	if len(nums) == 0 {
		return
	}

	m := len(nums) / 2
	n := TreeNode{Val: nums[m]}
	insertLeft(&n, nums[:m])
	insertRight(&n, nums[m+1:])
	root.Right = &n
}
