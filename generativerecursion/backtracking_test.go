package generativerecursion

import "testing"

func TestThreatening(t *testing.T) {
	tests := []struct {
		a, b QP
		ok   bool
	}{
		{QP{0, 0}, QP{0, 3}, true},
		{QP{0, 0}, QP{3, 0}, true},
		{QP{0, 0}, QP{1, 1}, true},
		{QP{0, 1}, QP{1, 0}, true},
		{QP{0, 0}, QP{1, 3}, false},
	}

	for i, tt := range tests {
		if ok := threatenging(tt.a, tt.b); ok != tt.ok {
			t.Errorf("test %d: want %v, got %v\n", i, tt.ok, ok)
		}
	}
}

func TestIsASolution(t *testing.T) {
	tests := map[string][]QP{
		"Queens4Solution1": Queen4Solution1,
		"Queens4Solution2": Queen4Solution2,
	}

	test := IsASolution(4)
	for name, tt := range tests {
		if !test(tt) {
			t.Errorf("%v is not a solution\n", name)
		}
	}
}

func TestNQueens(t *testing.T) {
	tests := []int{2, 3, 4}

	for i, n := range tests {
		s := NQueens(n)
		if n < 4 {
			if s != nil {
				t.Errorf("test %d: expected %d-queen problem has no solution, but it has: %v\n", i, n, s)
			}
		} else {
			if s == nil {
				t.Errorf("test %d: expected %d-queen problem has solution, but got no one\n", i, n)
			} else if !IsASolution(n)(s) {
				t.Errorf("test %d: bad %d-queen problem solution %v\n", i, n, s)
			}
		}
	}
}
