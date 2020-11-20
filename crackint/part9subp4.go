package crackint

// Trees and Graphs

type bstNode struct {
	v     int
	left  *bstNode
	right *bstNode
}

func insert(n *bstNode, v int) {
	if n == nil {
		panic("nil node")
	}

	if v <= n.v {
		if n.left != nil {
			insert(n.left, v)
			return
		}

		n.left = &bstNode{v: v}
		return
	}

	if n.right != nil {
		insert(n.right, v)
		return
	}

	n.right = &bstNode{v: v}
}

func inOrder(root *bstNode, f func(*bstNode)) {
	if root == nil {
		return
	}

	inOrder(root.left, f)
	f(root)
	inOrder(root.right, f)
}

func preOrder(root *bstNode, f func(*bstNode)) {
	if root == nil {
		return
	}

	f(root)
	preOrder(root.left, f)
	preOrder(root.right, f)
}

func postOrder(root *bstNode, f func(*bstNode)) {
	if root == nil {
		return
	}

	postOrder(root.left, f)
	postOrder(root.right, f)
	f(root)
}

// graph

func dfs(root int, graph map[int][]int, f func(int)) {
	if graph == nil {
		return
	}

	visited := map[int]struct{}{}

	dfsRecursive(root, graph, visited, f)
}

func dfsRecursive(root int, graph map[int][]int, visited map[int]struct{}, f func(int)) {
	if _, ok := visited[root]; !ok {
		visited[root] = struct{}{}
		f(root)
	}

	for _, n := range graph[root] {
		if _, ok := visited[n]; !ok {
			dfsRecursive(n, graph, visited, f)
		}
	}
}

func bfs(root int, graph map[int][]int, f func(int)) {
	if graph == nil {
		return
	}

	queue := []int{root}
	visited := map[int]struct{}{}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, s := range graph[n] {
			if _, ok := visited[s]; !ok {
				queue = append(queue, s)
			}
		}

		if _, ok := visited[n]; !ok {
			f(n)
			visited[n] = struct{}{}
		}
	}
}
