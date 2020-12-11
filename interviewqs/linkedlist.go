package interviewqs

// design a doubli linked list
// add("B", 0)
// add("D", 1)
// add("C", 1)
// add("E", 3)
// add("A", 0)
// print(linkedlist) -> A, B, C, D, E
type linkedlist struct {
	root *listnode
}

func (l *linkedlist) add(v string, pos int) {
	if l.root == nil {
		l.root = &listnode{v: v}
		return
	}

	l.addRecursive(l.root, v, 0, pos)
	if l.root.prev != nil {
		l.root = l.root.prev
	}
}

func (l *linkedlist) addRecursive(node *listnode, v string, cur, pos int) {
	if cur == pos {
		insert := &listnode{v: v, prev: node.prev, next: node}
		if node.prev != nil {
			node.prev.next = insert
		}
		node.prev = insert
		return
	}

	if node.next != nil {
		l.addRecursive(node.next, v, cur+1, pos)
		return
	}

	node.next = &listnode{v: v, prev: node}
}

type listnode struct {
	v          string
	prev, next *listnode
}
