package generativerecursion

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, d int
	}{
		{2, 3, 1},
		{6, 25, 1},
		{18, 24, 6},
		{101135853, 45014640, 177},
	}

	for i, tt := range tests {
		o := GCD(tt.a, tt.b)
		if o != tt.d {
			t.Errorf("test %d: want %d, got %d\n", i, tt.d, o)
		}
	}
}
