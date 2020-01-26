package generativerecursion

import "math"

// Queens is a number of queens in a chess table
var Queens = 4

// QP is a queen position,
// where Row and Col < Queens
type QP struct {
	Row, Col int
}

var (
	// Queen4Solution1 is a 4 queens problem solution
	Queen4Solution1 = []QP{QP{0, 2}, QP{1, 0}, QP{2, 3}, QP{3, 1}}

	// Queen4Solution2 is another 4 queens problem solution
	Queen4Solution2 = []QP{QP{0, 1}, QP{1, 3}, QP{2, 0}, QP{3, 2}}
)

// NQueens returns all solutions for n-queens problem if any
func NQueens(n int) []QP {
	return placeQueens(nil, n)
}

func threatenging(a, b QP) bool {
	return a.Row == b.Row || a.Col == b.Col || math.Abs(float64(a.Row-b.Row)) == math.Abs(float64(a.Col-b.Col))
}

// IsASolution returns a function to test is a given Queens Positions
// is a solution for n queens problem
func IsASolution(n int) func([]QP) bool {
	return func(s []QP) bool {
		if len(s) != n {
			return false
		}

		for i, qp1 := range s[:len(s)-1] {
			for _, qp2 := range s[i+1:] {
				if threatenging(qp1, qp2) {
					return false
				}
			}
		}
		return true
	}
}

func placeQueens(board []QP, n int) []QP {
	if len(board) == n {
		return board
	}

	r := len(board)
	pp := possibles(board, r, n)
	if pp == nil {
		return nil
	}

	for _, p := range pp {
		board = append(board, p)
		b := placeQueens(board, n)
		if b != nil {
			return b
		}
		board = board[:len(board)-1]
	}
	return nil
}

func possibles(board []QP, r, n int) []QP {
	p := make([]QP, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, QP{r, i})
	}

	return filterQPs(p, func(qp QP) bool {
		for _, q := range board {
			if threatenging(qp, q) {
				return false
			}
		}
		return true
	})
}

func filterQPs(qps []QP, f func(QP) bool) []QP {
	var filtered []QP
	for _, qp := range qps {
		if f(qp) {
			filtered = append(filtered, qp)
		}
	}
	return filtered
}
