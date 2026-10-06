package networktimedelay

import "container/heap"

// An entry in our to-do list: "reach `node` at `time`"
type Item struct {
	time int
	node int
}

// MinHeap sorted by time (smallest first)
type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func NetworkDelayTime(times [][]int, n int, k int) int {
	// 1. Build the map: graph[u] = list of [neighbor, travel_time]
	graph := make(map[int][][2]int)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		graph[u] = append(graph[u], [2]int{v, w})
	}

	// 2. Finalized shortest times
	dist := make(map[int]int)

	// 3. Min-heap, start at k with time 0
	h := &MinHeap{{0, k}}

	for h.Len() > 0 {
		cur := heap.Pop(h).(Item) // closest house first

		// Skip if already finalized (stale entry)
		if _, done := dist[cur.node]; done {
			continue
		}

		dist[cur.node] = cur.time // shortest time found!

		// Try every road leaving this house
		for _, edge := range graph[cur.node] {
			neighbor, w := edge[0], edge[1]
			if _, done := dist[neighbor]; !done {
				heap.Push(h, Item{cur.time + w, neighbor})
			}
		}
	}

	// 4. Did everyone hear it?
	if len(dist) != n {
		return -1
	}

	// 5. The slowest house decides the answer
	maxTime := 0
	for _, t := range dist {
		if t > maxTime {
			maxTime = t
		}
	}
	return maxTime
}
