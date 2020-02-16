package sortings

import "testing"

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
		}
	}
}

func equal(a1, a2 []Value) bool {
	if len(a1) == 0 && len(a2) == 0 {
		return true
	}

	if a1[0].IntValue != a2[0].IntValue || a1[0].StringValue != a2[0].StringValue {
		return false
	}

	return equal(a1[1:], a2[1:])
}
