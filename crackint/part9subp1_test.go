package crackint

import "testing"

func TestIsUniqueChars(t *testing.T) {
	tt := []struct {
		s  string
		ok bool
	}{
		{"aa", false},
		{"ab", true},
	}

	for _, tc := range tt {
		t.Run(tc.s, func(t *testing.T) {
			ok := isUniqueChars(tc.s)
			if ok != tc.ok {
				t.Fatalf("want %v, got %v", tc.ok, ok)
			}
		})
	}
}

func TestCheckPermutation(t *testing.T) {
	tt := []struct {
		a, b string
		ok   bool
	}{
		{"a", "a", true},
		{"a", "A", false},
		{"a", "aa", false},
		{"asdff", "fafds", true},
		{"", "", true},
	}

	for _, tc := range tt {
		t.Run(tc.a+"_"+tc.b, func(t *testing.T) {
			ok := checkPermutation(tc.a, tc.b)
			if ok != tc.ok {
				t.Fatalf("want %v, got %v", tc.ok, ok)
			}
		})
	}
}
