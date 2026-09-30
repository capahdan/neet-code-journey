package graphvalidtree

// ValidTree checks whether n noe (0..n-1) and the given undirected edges
// form a valid tree
//
// A graph of n nodes is a tree if
// 1. It has exactly n-1 edges (a tree with n nodes always has n-1 edges), and
// 2. It has no cycles (which, combined with n-1 edges, guarantee it's fully connected
//    fully connected as a single component).
//
// We use Union-Find : for each edge, if both endpoints already belong to the same set
// adding this edge would create a cycle -> not. a tree.

func ValidTree(n int, edges [][]int) bool {
	//  A tree mas has excatly n-1 edges.
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	// find returns the root of x's set, with path compression
	var find func(x int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // path compression halving
			x = parent[x]
		}

		return x
	}

	for _, e := range edges {
		rootA, rootB := find(e[0]), find(e[1])
		if rootA == rootB {
			// both nodes already connected -> this edge creates a cyle
			return false
		}
		parent[rootA] = rootB // union
	}

	// len(edges) == n-1 and no cycles found -> guaranteed fully connected.
	return true
}
