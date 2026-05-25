package laststoneweight

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	length := len(old)
	x := old[length-1]
	*m = old[:length-1]

	return x
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, st := range stones {
		heap.Push(h, st)
	}

	for h.Len() > 1 {
		firstLargest := heap.Pop(h).(int)
		secondLargest := heap.Pop(h).(int)

		if firstLargest != secondLargest {
			heap.Push(h, firstLargest-secondLargest)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return (*h)[0]
}
