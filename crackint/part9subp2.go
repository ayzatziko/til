package crackint

type node struct {
	v    int
	next *node
}

func makeLinkedList(ints ...int) *node {
	if ints == nil {
		return nil
	}

	head := &node{v: ints[0]}
	cur := head

	for _, v := range ints[1:] {
		cur.next = &node{v: v}
		cur = cur.next
	}

	return head
}

func doLinkedListRecursive(n *node, f func(*node)) {
	if n == nil {
		return
	}

	f(n)
	doLinkedListRecursive(n.next, f)
}

// 2.1 linked list dedup
func dedupV1(head *node) {
	if head == nil {
		return
	}

	dups := map[int]struct{}{
		head.v: {},
	}

	prev, cur := head, head.next
	for ; cur != nil; prev, cur = cur, cur.next {
		if _, ok := dups[cur.v]; ok {
			prev.next = cur.next
			cur.next = nil
		}

		dups[cur.v] = struct{}{}
	}
}

func kthToLast(head *node, i int) *node {
	n, _ := kthToLastRecursive(head, i)
	return n
}

func kthToLastRecursive(n *node, i int) (*node, int) {
	if n == nil {
		return nil, -1
	}

	x, xn := kthToLastRecursive(n.next, i)
	if x != nil {
		return x, xn + 1
	}

	if xn+1 == i {
		return n, xn + 1
	}

	return nil, xn + 1
}
