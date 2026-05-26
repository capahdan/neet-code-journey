package kthlargestelementinanarray

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() any {
	old := *m
	length := len(old)

	x := old[length-1]
	*m = old[:length-1]

	return x
}

func findKthLargest(nums []int, k int) int {

	h := &MinHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
		fmt.Println(h)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
