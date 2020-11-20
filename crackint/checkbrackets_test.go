package crackint

import "testing"

func TestCheckBrackets(t *testing.T) {
	tt := []struct {
		s  string
		ok bool
	}{
		{"", true},
		{"(", false},
		{"{{", false},
		{"({}[[]]{})", true},
	}

	for _, tc := range tt {
		name := tc.s
		if name == "" {
			name = "empty"
		}

		t.Run(name, func(t *testing.T) {
			ok := checkBrackets(tc.s)
			if ok != tc.ok {
				t.Fatalf("want %v, got %v", tc.ok, ok)
			}
		})
	}
}
