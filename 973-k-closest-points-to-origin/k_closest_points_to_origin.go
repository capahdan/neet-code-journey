package kclosestpointstoorigin

import "container/heap"

type MaxHeap [][]int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	distI := m[i][0]*m[i][0] + m[i][1]*m[i][1]
	distJ := m[j][0]*m[j][0] + m[j][1]*m[j][1]

	return distI > distJ

}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *MaxHeap) Pop() any {

	old := *m
	length := len(old)

	x := old[length-1]
	*m = old[:length-1]

	return x
}

func kClosest(points [][]int, k int) [][]int {

	h := &MaxHeap{}
	heap.Init(h)

	for _, point := range points {
		heap.Push(h, point)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return [][]int(*h)
}
