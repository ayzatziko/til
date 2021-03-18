package lexer

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		input  string
		result int
	}{
		{"5", 5},
		{"", 0},
		{"23+32", 55},
		{"52-29", 23},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%d", tt.input, tt.result), func(t *testing.T) {
			res := eval(tt.input)

			if res != tt.result {
				t.Fatalf("expected %s = %d, got %d", tt.input, tt.result, res)
			}
		})
	}
}
