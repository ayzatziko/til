package abstractions

func Map(slice []int, f func(int) int) []int {
	a := make([]int, 0, len(slice))
	for _, v := range slice {
		a = append(a, f(v))
	}
	return a
}

func Filter(slice []int, f func(int) bool) []int {
	var a []int
	for _, v := range slice {
		if f(v) {
			a = append(a, v)
		}
	}
	return a
}

func Sort(slice []int, f func(int, int) bool) []int {
	if len(slice) == 0 {
		return nil
	}

	return insert(Sort(slice[1:], f), slice[0], f)
}

func insert(a []int, v int, f func(int, int) bool) []int {
	if len(a) == 0 {
		return []int{v}
	}

	sorted := make([]int, 0, len(a)+1)

	if f(v, a[0]) {
		sorted = append(sorted, a[0])
		return append(sorted, insert(a[1:], v, f)...)
	}

	sorted = append(sorted, v)
	return append(sorted, a...)
}

func AndMap(slice []int, f func(int) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}

	return true
}
