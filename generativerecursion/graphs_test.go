package generativerecursion

import (
	"reflect"
	"testing"
)

func TestGraphs(t *testing.T) {
	g := Graph{"A": []string{"B", "E"}, "B": []string{"E", "F"}, "C": []string{"D"}, "D": nil, "E": []string{"C", "F"}, "F": []string{"D", "G"}, "G": nil}

	tests := []struct {
		s, d string
		p    []string
	}{
		{"C", "D", []string{"C", "D"}},
		{"E", "D", []string{"E", "C", "D"}},
		{"A", "G", []string{"A", "B", "E", "F", "G"}},
		{"C", "G", nil},
	}

	for i, tt := range tests {
		p := FindPath(tt.s, tt.d, g)
		if !reflect.DeepEqual(tt.p, p) {
			t.Errorf("test %d: want %v, got %v\n", i, tt.p, p)
		}
	}
}
