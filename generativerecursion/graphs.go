package generativerecursion

// Graph is a directed acyclic graph (DAG)
type Graph map[string][]string

// FindPath returns a path from s node to d node if any,
// ohterwise returns nil
//
// nodes: A, B, C, D, E, F, G
// edges: {A:[B,E], B: [E, F], C: [D], D: [], E: [C, F], F: [D, G], G: []}
// FindPath("C", "D", g) -> [C, D]
// FindPath("E", "D", g) -> [E, C, D]
// FindPath("A", "G", g) -> [A, B, E, F, G]
// FindPath("C", "G", g) -> nil
func FindPath(s, d string, g Graph) []string {
	if s == d {
		return []string{s}
	}

	neighbors := g[s]
	var p []string
	for _, n := range neighbors {
		p = FindPath(n, d, g)
		if p != nil {
			return append([]string{s}, p...)
		}
	}

	return nil
}
