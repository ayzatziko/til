// problem #118 (https://leetcode.com/problems/pascals-triangle/)
package leetcode

func PascalsTriangle(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	result := [][]int{[]int{1}}
	for i := 0; i < numRows-1; i++ {
		result = append(result, next(result[i]))
	}
	return result
}

func next(prev []int) []int {
	n := make([]int, 0, len(prev)+1)
	n = append(n, 1)
	for i := 0; i < len(prev)-1; i++ {
		n = append(n, prev[i]+prev[i+1])
	}
	n = append(n, 1)
	return n
}
