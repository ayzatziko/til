// problem #103 (https://leetcode.com/problems/binary-tree-zigzag-level-order-traversa)
package leetcode

func ZigzagLevelOrder(root *TreeNode) [][]int {
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

	for i := range result {
		if i%2 == 1 {
			result[i] = reverse(result[i])
		}
	}

	return result
}

func reverse(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}
