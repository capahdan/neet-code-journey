package mincosttoconnectallnodes

import "container/heap"

type item struct {
	x        int
	y        int
	distance int
}

type minHeap []item

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].distance < m[j].distance }
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

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	graph := make(map[[2]int][]item)

	for i := 0; i < n; i++ {
		key := [2]int{points[i][0], points[i][1]}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dist := mutlak(points[i][0]-points[j][0]) + mutlak(points[i][1]-points[j][1])
			graph[key] = append(graph[key], item{
				x:        points[j][0],
				y:        points[j][1],
				distance: dist,
			})
		}
	}

	var result int
	visited := make(map[[2]int]bool)
	h := &minHeap{}

	heap.Push(h, item{
		x:        points[0][0],
		y:        points[0][1],
		distance: 0,
	})

	for h.Len() > 0 {
		curr := heap.Pop(h).(item)
		currKey := [2]int{curr.x, curr.y}

		if visited[currKey] {
			continue
		}
		visited[currKey] = true
		result += curr.distance
		for _, next := range graph[currKey] {
			nextKey := [2]int{next.x, next.y}
			if !visited[nextKey] {
				heap.Push(h, next)
			}
		}
	}

	return result
}

func mutlak(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
