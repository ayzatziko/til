package sortings

type Value struct {
	IntValue    int
	StringValue string
}

func MergeSort(arr []Value) []Value {
<<<<<<< HEAD
	if arr == nil {
=======
	if arr == nil || len(arr) == 0 {
>>>>>>> 4cbc8d2... added package sortings and sort implementations
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
<<<<<<< HEAD
		if compare(a1[0], a2[0]) {
=======
		if sortFunc(a1[0], a2[0]) {
>>>>>>> 4cbc8d2... added package sortings and sort implementations
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

<<<<<<< HEAD
func compare(a, b Value) bool {
	return a.IntValue <= b.IntValue
=======
func QuickSort(arr []Value) []Value {
	if len(arr) == 0 {
		return nil
	}

	p := arr[0]
	less := QuickSort(filterValue(arr[1:], func(v Value) bool {
		return v.IntValue <= p.IntValue
	}))
	more := QuickSort(filterValue(arr[1:], func(e Value) bool {
		return e.IntValue > p.IntValue
	}))
	sarr := make([]Value, 0, len(less)+len(more)+1)
	sarr = append(sarr, less...)
	sarr = append(sarr, p)
	sarr = append(sarr, more...)

	return sarr
}

func filterValue(arr []Value, f func(Value) bool) []Value {
	var farr []Value
	for _, v := range arr {
		if f(v) {
			farr = append(farr, v)
		}
	}
	return farr
}

func HeapSort(arr []Value) []Value {
	return nil
}

func BubbleSort(arr []Value) []Value {
	if len(arr) == 0 {
		return nil
	}

	sorted := BubbleSort(arr[1:])
	return insert(sorted, arr[0])
}

func insert(arr []Value, v Value) []Value {
	if len(arr) == 0 {
		return []Value{v}
	}

	if !sortFunc(arr[0], v) {
		sarr := make([]Value, 0, len(arr)+1)
		sarr = append(sarr, v)
		return append(sarr, arr...)
	}

	return append([]Value{arr[0]}, insert(arr[1:], v)...)
}

func areSorted(arr []Value) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}

	if !sortFunc(arr[0], arr[1]) {
		return false
	}

	return areSorted(arr[1:])
}

func sortFunc(v0, v1 Value) bool {
	return v0.IntValue <= v1.IntValue
>>>>>>> 4cbc8d2... added package sortings and sort implementations
}
