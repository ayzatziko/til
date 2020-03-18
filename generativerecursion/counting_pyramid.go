package generativerecursion

func CountPyramid(levels int) [][]int {
	result := make([][]int, levels)
	result[0] = []int{1}

	for i := 0; i < levels-1; i++ {
		result[i+1] = countLevel(result[i])
	}

	return result
}

func countLevel(arr []int) []int {
	var (
		result []int
	)

	for len(arr) > 0 {
		count := firstInARow(arr)
		result = append(result, count, arr[0])
		arr = arr[count:]
	}

	return result
}

func firstInARow(arr []int) int {
	v := arr[0]
	for i, n := range arr {
		if n != v {
			return i
		}
	}

	return len(arr)
}
