package generativerecursion

// File contains Gauss elimination algorithm

// SOE is a non empty matrix
// SOE = system of equations like
// 	2x + 2y + 3z  = 10
// 	2x + 5y + 12z = 31
// 	4x + y  - 2z  = 1
//
// 	data example: [][]int{{2,2,3,10}, {2,5,12,31}, {4,1,-2,1}}
type SOE []Equation

// TSOE is a triangular SOE
// such that the Equations are of decreasing length:
// 	n, n-1, ..., 2
type TSOE SOE

// Equation is an array [a0 a1 ... an b] of at least 2 numbers, where
// a[i]s are the left-hand-side and b is right-hand-side
type Equation []int

func Triangulate(soe SOE) TSOE {
	if len(soe) == 1 {
		return TSOE(soe)
	}

	substracted := mapSOE(soe[1:], func(e Equation) Equation {
		return substract(soe[0], e, e[0], soe[0][0])
	})

	firsts := mapSOE(substracted, func(e Equation) Equation { return e[1:] })
	if firsts[0][0] == 0 {
		i, ok := findEquation(firsts, func(e Equation) bool { return e[0] != 0 })
		if !ok {
			return nil
		}
		firsts[0], firsts[i] = firsts[i], firsts[0]
	}

	return append(TSOE{soe[0]}, mapSOE(SOE(Triangulate(firsts)), func(e Equation) Equation {
		eq := make(Equation, 0, len(e)+1)
		eq = append(eq, 0)
		return append(eq, e...)
	})...)
}

// Solution is a SOE solution
type Solution []int

func CheckSolution(soe SOE, sol Solution) bool {
	res := func(e Equation) int {
		r := 0
		for i, v := range lhs(e) {
			r += sol[i] * v
		}
		return r
	}

	for _, eq := range soe {
		if res(eq) != rhs(eq) {
			return false
		}
	}
	return true
}

func Solve(tsoe TSOE) Solution {
	if len(tsoe) == 1 {
		return Solution{tsoe[0][1] / tsoe[0][0]}
	}
	firsts := mapTSOE(tsoe[1:], func(e Equation) Equation { return e[1:] })
	sol := Solve(firsts)

	lst := len(tsoe[0]) - 1
	for i := range sol {
		tsoe[0][lst] -= tsoe[0][i+1] * sol[i]
	}
	return append(Solution{tsoe[0][lst] / tsoe[0][0]}, sol...)
}

// whishlist below

func substract(e1, e2 Equation, m1, m2 int) Equation {
	if len(e1) == 1 {
		return Equation{e1[0]*m1 - e2[0]*m2}
	}
	return append(Equation{e1[0]*m1 - e2[0]*m2}, substract(e1[1:], e2[1:], m1, m2)...)
}

func mapSOE(soe SOE, f func(Equation) Equation) SOE {
	nsoe := make(SOE, 0, len(soe))
	for _, e := range soe {
		nsoe = append(nsoe, f(e))
	}
	return nsoe
}

func mapTSOE(tsoe TSOE, f func(Equation) Equation) TSOE {
	ntsoe := make(TSOE, 0, len(tsoe))
	for _, e := range tsoe {
		ntsoe = append(ntsoe, f(e))
	}
	return ntsoe
}

func lhs(e Equation) []int {
	vars := make([]int, 0, len(e)-1)
	for _, v := range e[:len(e)-1] {
		vars = append(vars, v)
	}
	return vars
}

func rhs(e Equation) int { return e[len(e)-1] }

func mapEquation(e Equation, f func(int) int) Equation {
	eq := make(Equation, 0, len(e))
	for _, v := range e {
		eq = append(eq, f(v))
	}
	return eq
}

func findEquation(soe SOE, f func(Equation) bool) (int, bool) {
	for i, e := range soe {
		if f(e) {
			return i, true
		}
	}
	return -1, false
}
