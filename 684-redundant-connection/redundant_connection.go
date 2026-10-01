package redundantconnection

func FindRedundantConnection(edges [][]int) []int {
	result := [][]int{}
	par := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)

	for i := range edges {
		par[i+1] = i + 1
		rank[i+1] = 1
	}

	var find func(n int) int
	find = func(n int) int {
		res := n
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
			par[p2] = p1
			rank[p2] += rank[p1]
		} else {
			par[p1] = p2
			rank[p1] += rank[p2]
		}

		return 1
	}

	for _, e := range edges {
		temp := union(e[0], e[1])
		if temp == 0 {
			result = append(result, e)
		}
	}

	return result[len(result)-1]
}
