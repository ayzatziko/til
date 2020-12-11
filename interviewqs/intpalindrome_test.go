package interviewqs

import (
	"fmt"
	"testing"
)

func TestIsIntPalindrome(t *testing.T) {
	tt := []struct {
		v  int
		ok bool
	}{
		{12321, true}, {1221, true}, {1232, false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d[%v]", tc.v, tc.ok), func(t *testing.T) {
			if ok := isIntPalindrome(tc.v); ok != tc.ok {
				t.Fatalf("isIntPalindrome(%d) returned %v, expected %v to be returned", tc.v, ok, tc.ok)
			}
		})
	}
}
