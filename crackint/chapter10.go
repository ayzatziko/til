package crackint

// Sorting

func mergeSort(arr []int) {
	helper := make([]int, len(arr))
	mergeSortRecursiveHelper(arr, helper)
}

func mergeSortRecursiveHelper(arr, helper []int) {
	if len(arr) > 1 {
		arrLen := len(arr)
		mid := arrLen / 2
		mergeSortRecursiveHelper(arr[:mid], helper[:mid])
		mergeSortRecursiveHelper(arr[mid:], helper[mid:])
		merge(arr, helper, mid)
	}
}

func merge(arr, helper []int, mid int) {
	copy(helper, arr)

	cur := 0
	leftH := 0
	rightH := mid
	max := len(helper)
	for leftH < mid && rightH < max {
		if helper[leftH] < helper[rightH] {
			arr[cur] = helper[leftH]
			leftH++
		} else {
			arr[cur] = helper[rightH]
			rightH++
		}
		cur++
	}

	for leftH < mid {
		arr[cur] = helper[leftH]
		cur++
		leftH++
	}
}
