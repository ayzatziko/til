package sorting

// https://www.hackerearth.com/ru/practice/algorithms/sorting/selection-sort/tutorial/

func SelectionSortN(arr []int, steps int) []int {
	selectionSortN(arr, steps)
	return arr
}

func selectionSortN(arr []int, steps int) {
	if steps == 0 {
		return
	}

	i := min(arr)
	arr[0], arr[i] = arr[i], arr[0]
	selectionSortN(arr[1:], steps-1)
}

func min(arr []int) int {
	m := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[m] {
			m = i
		}
	}
	return m
}
