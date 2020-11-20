package crackint

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	tt := []struct {
		ints []int
		out  []int
	}{
		{[]int{5, 3, 1, 4, 7, 6, 8}, []int{5, 3, 1, 4, 7, 6, 8}},
		{nil, nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.ints, tc.out), func(t *testing.T) {
			var acc []int

			var r *bstNode
			if tc.ints != nil {
				acc = make([]int, 0, len(tc.ints))
				r = &bstNode{v: tc.ints[0]}

				for _, v := range tc.ints[1:] {
					insert(r, v)
				}
			}

			preOrder(r, func(n *bstNode) {
				acc = append(acc, n.v)
			})

			if len(acc) != len(tc.out) {
				t.Fatalf("expected %d elements, got %d", len(tc.out), len(acc))
			}

			for i, v := range tc.out {
				if v != acc[i] {
					t.Fatalf("expected out[%d] is %d, got %d", i, v, acc[i])
				}
			}
		})
	}
}

func TestInOrder(t *testing.T) {
	tt := []struct {
		ints []int
		out  []int
	}{
		{[]int{5, 3, 1, 4, 7, 6, 8}, []int{1, 3, 4, 5, 6, 7, 8}},
		{nil, nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.ints, tc.out), func(t *testing.T) {
			var acc []int

			var r *bstNode
			if tc.ints != nil {
				acc = make([]int, 0, len(tc.ints))
				r = &bstNode{v: tc.ints[0]}

				for _, v := range tc.ints[1:] {
					insert(r, v)
				}
			}

			inOrder(r, func(n *bstNode) {
				acc = append(acc, n.v)
			})

			if len(acc) != len(tc.out) {
				t.Fatalf("expected %d elements, got %d", len(tc.out), len(acc))
			}

			for i, v := range tc.out {
				if v != acc[i] {
					t.Fatalf("expected out[%d] is %d, got %d", i, v, acc[i])
				}
			}
		})
	}
}

func TestPostOrder(t *testing.T) {
	tt := []struct {
		ints []int
		out  []int
	}{
		{[]int{5, 3, 1, 4, 7, 6, 8}, []int{1, 4, 3, 6, 8, 7, 5}},
		{nil, nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.ints, tc.out), func(t *testing.T) {
			var acc []int

			var r *bstNode
			if tc.ints != nil {
				acc = make([]int, 0, len(tc.ints))
				r = &bstNode{v: tc.ints[0]}

				for _, v := range tc.ints[1:] {
					insert(r, v)
				}
			}

			postOrder(r, func(n *bstNode) {
				acc = append(acc, n.v)
			})

			if len(acc) != len(tc.out) {
				t.Fatalf("expected %d elements, got %d", len(tc.out), len(acc))
			}

			for i, v := range tc.out {
				if v != acc[i] {
					t.Fatalf("expected out is %v, got %v", tc.out, acc)
				}
			}
		})
	}
}

func TestDFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 4, 5}, 1: {3, 4}, 2: {1}, 3: {2, 4}, 4: {}, 5: {},
	}

	acc := []int{}
	dfs(0, graph, func(v int) {
		acc = append(acc, v)
	})

	res := []int{0, 1, 3, 2, 4, 5}
	if len(res) != len(acc) {
		t.Fatalf("expected %d ints, got %d", len(res), len(acc))
	}

	for i, v := range res {
		if v != acc[i] {
			t.Fatalf("expected %v, got %v", res, acc)
		}
	}
}

func TestBFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 4, 5}, 1: {3, 4}, 2: {1}, 3: {2, 4}, 4: {}, 5: {},
	}

	acc := []int{}
	bfs(0, graph, func(v int) {
		acc = append(acc, v)
	})

	res := []int{0, 1, 4, 5, 3, 2}
	if len(res) != len(acc) {
		t.Fatalf("expected %d ints, got %d", len(res), len(acc))
	}

	for i, v := range res {
		if v != acc[i] {
			t.Fatalf("expected %v, got %v", res, acc)
		}
	}
}
