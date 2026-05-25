package kthlargestelementinstream

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[:n-1]

	return x
}

type KthLargest struct {
	k int
	h *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := KthLargest{k: k, h: h}

	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.h, val)

	if kl.h.Len() > kl.k {
		heap.Pop(kl.h)
	}

	return (*kl.h)[0]
}
