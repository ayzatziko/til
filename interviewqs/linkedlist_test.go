package interviewqs

import "testing"

func TestLinkedList(t *testing.T) {
	type add struct {
		v   string
		pos int
	}

	tt := []struct {
		adds []add
		res  string
	}{
		{[]add{{"B", 0}, {"D", 1}, {"C", 1}, {"E", 3}, {"A", 0}}, "A, B, C, D, E"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			l := linkedlist{}
			for _, add := range tc.adds {
				l.add(add.v, add.pos)
			}

			s := ""
			for n := l.root; n != nil; n = n.next {
				if n.next == nil {
					s += n.v
				} else {
					s += n.v + ", "
				}
			}

			if tc.res != s {
				t.Fatalf("expected %s, got %s", tc.res, s)
			}
		})
	}
}
