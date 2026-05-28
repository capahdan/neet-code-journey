package taskscheduler

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
	l := len(old)

	x := old[l-1]
	*m = old[:l-1]
	return x
}

type Item struct {
	count   int
	readyAt int
}

func leastInterval(tasks []byte, n int) int {
	freqMap := make(map[byte]int)
	for _, t := range tasks {
		freqMap[t]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range freqMap {
		heap.Push(h, val)
	}

	queue := []Item{}
	time := 0

	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			cnt := heap.Pop(h).(int) - 1
			if cnt > 0 {
				queue = append(queue, Item{count: cnt, readyAt: time + n})
			}
		}

		if len(queue) > 0 && queue[0].readyAt == time {
			heap.Push(h, queue[0].count)
			queue = queue[1:]
		}

	}

	return time
}
