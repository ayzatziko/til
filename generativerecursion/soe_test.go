package generativerecursion

import (
	"reflect"
	"testing"
)

func TestSOECheckSolution(t *testing.T) {
	tests := []struct {
		soe SOE
		sol Solution
	}{
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{2, 5, 12, 31},
				Equation{4, 1, -2, 1},
			}, Solution{1, 1, 2},
		},
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{0, 3, 9, 21},
				Equation{0, 0, 1, 2},
			}, Solution{1, 1, 2},
		},
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{0, 3, 9, 21},
				Equation{0, -3, -8, -19},
			}, Solution{1, 1, 2},
		},
	}

	for i, tt := range tests {
		if !CheckSolution(tt.soe, tt.sol) {
			t.Errorf("test %d: invalid solution\n", i)
		}
	}
}

func TestTriangulate(t *testing.T) {

	tests := []struct {
		in  SOE
		out TSOE
	}{
		{
			SOE{Equation{2, 2, 3, 10}, Equation{2, 5, 12, 31}, Equation{4, 1, -2, 1}},
			TSOE{Equation{2, 2, 3, 10}, Equation{0, -6, -18, -42}, Equation{0, 0, -12, -24}},
		},
		{
			SOE{Equation{2, 3, 3, 8}, Equation{2, 3, -2, 3}, Equation{4, -2, 2, 4}},
			TSOE{Equation{2, 3, 3, 8}, Equation{0, 16, 8, 24}, Equation{0, 0, -160, -160}},
		},
	}

	for i, tt := range tests {
		out := Triangulate(tt.in)
		if !equalTSOE(tt.out, out) {
			t.Errorf("test %d: want %v, got %v\n", i, tt.out, out)
		}
	}
}

func equalTSOE(t1, t2 TSOE) bool {
	if len(t1) == 0 {
		return true
	}

	for i := range t1 {
		for k := range t1[i] {
			if t1[i][k] != t2[i][k] {
				return false
			}
		}
	}

	return true
}

func TestSolve(t *testing.T) {
	tests := []struct {
		soe SOE
		sol Solution
	}{
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{2, 5, 12, 31},
				Equation{4, 1, -2, 1},
			}, Solution{1, 1, 2},
		},
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{0, 3, 9, 21},
				Equation{0, 0, 1, 2},
			}, Solution{1, 1, 2},
		},
		{
			SOE{
				Equation{2, 2, 3, 10},
				Equation{0, 3, 9, 21},
				Equation{0, -3, -8, -19},
			}, Solution{1, 1, 2},
		},
	}

	for i, tt := range tests {
		s := Solve(Triangulate(tt.soe))
		if !reflect.DeepEqual(tt.sol, s) {
			t.Errorf("test %d: want %v, got %v\n", i, tt.sol, s)
		}
	}
}
