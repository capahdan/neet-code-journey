package handofstraight

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand) % groupSize

	if n != 0 {
		return false
	} else {
		numMap := make(map[int]int)
		h := &minHeap{}
		for _, v := range hand {
			if _, ok := numMap[v]; !ok {
				heap.Push(h, v)
			}
			numMap[v]++
		}

		for h.Len() > 0 {
			start := (*h)[0]
			for v := start; v < start+groupSize; v++ {
				if numMap[v] == 0 {
					return false
				}
				numMap[v]--
				if numMap[v] == 0 {
					// v is exhausted, so it must be the current minimum —
					// otherwise something smaller is still unmatched.
					if (*h)[0] != v {
						return false
					}
					heap.Pop(h)
				}
			}
		}
	}

	return true
}

func isNStraightHand2(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	count := make(map[int]int)
	for _, v := range hand {
		count[v]++
	}

	key := make([]int, len(count))
	for k := range count {
		key = append(key, k)
	}
	sort.Ints(key)

	for _, k := range key {
		need := count[k]
		if need == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if count[k+i] < need {
				return false
			}
			count[k+i] -= need
		}
	}

	return true
}
