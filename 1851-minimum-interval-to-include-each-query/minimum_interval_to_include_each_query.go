package minimumintervaltoincludeeachquery

import (
	"container/heap"
	"sort"
)

type item struct {
	size  int
	right int
}

type MinHeap []item

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(item))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].size < m[j].size
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func minInterval(intervals [][]int, queries []int) []int {

	answer := make([]int, len(queries))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	order := make([]int, len(queries))
	for i := range queries {
		order[i] = i
	}

	sort.Slice(order, func(a, b int) bool {
		return queries[order[a]] < queries[order[b]]
	})

	h := &MinHeap{}
	i := 0

	for _, qi := range order {
		d := queries[qi]

		for i < len(queries) && intervals[i][0] <= d {
			l, r := intervals[i][0], intervals[i][1]
			heap.Push(h, item{size: r - l + 1, right: r})
			i++
		}

		for h.Len() > 0 && (*h)[0].right < d {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			answer[qi] = (*h)[0].size
		} else {
			answer[qi] = -1
		}

	}

	return answer
}
