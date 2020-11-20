package crackint

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMakeLinkedList(t *testing.T) {
	var ints []int
	f := func(n *node) {
		ints = append(ints, n.v)
	}

	tt := []struct {
		ints []int
	}{
		{nil}, {[]int{1}}, {[]int{1, 2, 3}},
	}

	for _, tc := range tt {
		t.Run(strconv.Itoa(len(tc.ints)), func(t *testing.T) {
			n := makeLinkedList(tc.ints...)
			ints = nil
			doLinkedListRecursive(n, f)
			if len(tc.ints) != len(ints) {
				t.Fatalf("expected %d ints, got %d", len(tc.ints), len(ints))
			}
			for i := range tc.ints {
				if tc.ints[i] != ints[i] {
					t.Fatalf("expects ints[%d] is %d, got %d", i, tc.ints[i], ints[i])
				}
			}
		})
	}
}

func TestDedup(t *testing.T) {
	dedupFuncs := []func(*node){
		dedupV1,
	}

	tt := []struct {
		in, out []int
	}{
		{nil, nil},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 1}, []int{1, 2}},
	}

	for _, tc := range tt {
		for i, dedup := range dedupFuncs {
			t.Run(fmt.Sprintf("%d:%v->%v", i, tc.in, tc.out), func(t *testing.T) {
				ll := makeLinkedList(tc.in...)
				dedup(ll)
				var ints []int
				doLinkedListRecursive(ll, func(n *node) { ints = append(ints, n.v) })
				if len(ints) != len(tc.out) {
					t.Fatalf("expected %d ints, got %d", len(tc.out), len(ints))
				}

				for x := range ints {
					if ints[x] != tc.out[x] {
						t.Fatalf("expected ints[%d] is %d, got %d", x, tc.out[x], ints[x])
					}
				}
			})
		}
	}
}

func TestKthToLast(t *testing.T) {
	tt := []struct {
		in   []int
		i    int
		want *node
	}{
		{nil, 0, nil},
		{[]int{1}, 0, &node{v: 1}},
		{[]int{1}, -2, nil},
		{[]int{1}, 5, nil},
		{[]int{1, 2, 3}, 1, &node{v: 2}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%+v", tc.in, tc.want), func(t *testing.T) {
			n := kthToLast(makeLinkedList(tc.in...), tc.i)
			if (n == nil) != (tc.want == nil) {
				t.Fatalf("unexpected result: want %v, got %v", tc.want, n)
			}

			if n != nil && n.v != tc.want.v {
				t.Fatalf("want %d got %d", tc.want.v, n.v)
			}
		})
	}
}
