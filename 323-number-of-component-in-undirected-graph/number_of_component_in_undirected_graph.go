package numberofcomponentinundirectedgraph

func CountComponents(n int, edges [][]int) int {
	par := make([]int, n)
	rank := make([]int, n)

	for i := range n {
		par[i] = i
		rank[i] = 1
	}

	var find func(i int) int
	find = func(i int) int {
		res := i
		for res != par[res] {
			par[res] = par[par[res]]
			res = par[res]
		}

		return res
	}

	var union func(n1, n2 int) int
	union = func(n1, n2 int) int {
		p1, p2 := find(n1), find(n2)

		if p1 == p2 {
			return 0
		}

		if rank[p2] > rank[p1] {
			par[p1] = p2
			rank[p2] += 1
		} else {
			par[p2] = p1
			rank[p1] += 1
		}

		return 1
	}

	res := n

	for _, e := range edges {
		res -= union(e[0], e[1])
	}

	return res
}
