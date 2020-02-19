package sortings

import "testing"

<<<<<<< HEAD
func TestSort(t *testing.T) {
	tests := []struct {
		in   []Value
		want []Value
	}{
		{nil, nil},
		{[]Value{Value{3, ""}, Value{2, ""}, Value{1, ""}}, []Value{Value{1, ""}, Value{2, ""}, Value{3, ""}}},
	}

	for i, tt := range tests {
		got := MergeSort(tt.in)
		if !equal(tt.want, got) {
			t.Errorf("test %d: want %v, got %v\n", i, tt.want, got)
=======
func TestSorting(t *testing.T) {
	testValues := [][]Value{
		nil,
		[]Value{Value{1, "1"}, Value{3, "3"}, Value{100, "100"}, Value{-2, "-2"}, Value{0, "0"}, Value{-123123, "123123"}, Value{20, "20"}, Value{57, "57"}},
	}
	sortFuncs := map[string]func([]Value) []Value{
		"quick sort":  QuickSort,
		"bubble sort": BubbleSort,
		"merge sort":  MergeSort,
		"heap sort":   HeapSort,
	}

	for sortType, sort := range sortFuncs {
		for _, ee := range testValues {
			var tee []Value
			if len(ee) > 0 {
				tee = make([]Value, len(ee))
			}
			copy(tee, ee)
			got := sort(tee)
			if !areSorted(got) {
				t.Errorf("%s: not sorted\n", sortType)
			}
>>>>>>> 4cbc8d2... added package sortings and sort implementations
		}
	}
}

<<<<<<< HEAD
func equal(a1, a2 []Value) bool {
	if len(a1) == 0 && len(a2) == 0 {
		return true
	}

	if a1[0].IntValue != a2[0].IntValue || a1[0].StringValue != a2[0].StringValue {
		return false
	}

	return equal(a1[1:], a2[1:])
=======
func TestSorted(t *testing.T) {
	tests := []struct {
		ee     []Value
		sorted bool
	}{
		{ee: []Value{Value{0, "0"}, Value{1, "1"}, Value{2, "2"}, Value{2, "2"}}, sorted: true},
		{ee: nil, sorted: true},
		{ee: []Value{Value{1, "1"}, Value{0, "0"}}, sorted: false},
	}

	for i, tt := range tests {
		if sorted := areSorted(tt.ee); tt.sorted != sorted {
			t.Errorf("%d: expected sorted %v, got %v\n", i, tt.sorted, sorted)
		}
	}
>>>>>>> 4cbc8d2... added package sortings and sort implementations
}
