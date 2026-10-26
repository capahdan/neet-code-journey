package cheapestflightwithinkstop

type item struct {
	destination int
	price       int
	k           int
}

type minHeap []item

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].price < m[j].price }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(item))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
// 	graph := make([][]item, n)
// 	for _, f := range flights {
// 		graph[f[0]] = append(graph[f[0]], item{destination: f[1], price: f[2]})
// 	}

// 	// minStops[v] = fewest edges used in any popped state reaching v
// 	minStops := make([]int, n)
// 	for i := range minStops {
// 		minStops[i] = 1 << 30
// 	}

// 	h := &minHeap{}
// 	heap.Push(h, item{destination: src, price: 0, k: 0})

// 	for h.Len() > 0 {
// 		cur := heap.Pop(h).(item)

// 		if cur.destination == dst {
// 			return cur.price
// 		}
// 		if cur.k > k { // k stops = k+1 edges; k here counts edges, so allow up to k+1
// 			continue
// 		}
// 		// Popped in price order, so an earlier pop was cheaper.
// 		// If it also used <= edges, this state is dominated.
// 		if cur.k >= minStops[cur.destination] {
// 			continue
// 		}
// 		minStops[cur.destination] = cur.k

// 		for _, nb := range graph[cur.destination] {
// 			heap.Push(h, item{
// 				destination: nb.destination,
// 				price:       cur.price + nb.price,
// 				k:           cur.k + 1,
// 			})
// 		}
// 	}
// 	return -1
// }

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1 << 30

	prices := make([]int, n)
	for i := range prices {
		prices[i] = inf
	}
	prices[src] = 0

	// k stops = at most k+1 edges
	for i := 0; i <= k; i++ {
		next := make([]int, n)
		copy(next, prices) // copy so each round only uses the previous round's values

		for _, f := range flights {
			from, to, cost := f[0], f[1], f[2]
			if prices[from] == inf {
				continue
			}
			if prices[from]+cost < next[to] {
				next[to] = prices[from] + cost
			}
		}
		prices = next
	}

	if prices[dst] == inf {
		return -1
	}
	return prices[dst]
}
