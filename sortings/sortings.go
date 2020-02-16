package sortings

type Value struct {
	IntValue    int
	StringValue string
}

func MergeSort(arr []Value) []Value {
	if arr == nil {
		return nil
	}

	arrl := len(arr)
	mid := arrl / 2

	if arrl == 1 {
		return arr
	}

	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func merge(a1, a2 []Value) []Value {
	var (
		sorted = make([]Value, 0, len(a1)+len(a2))
	)
	for len(a1) > 0 && len(a2) > 0 {
		if compare(a1[0], a2[0]) {
			sorted = append(sorted, a1[0])
			a1 = a1[1:]
		} else {
			sorted = append(sorted, a2[0])
			a2 = a2[1:]
		}
	}
	sorted = append(sorted, a1...)
	sorted = append(sorted, a2...)
	return sorted
}

func compare(a, b Value) bool {
	return a.IntValue <= b.IntValue
}
