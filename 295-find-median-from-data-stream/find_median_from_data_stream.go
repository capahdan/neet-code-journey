package findmedianfromdatastream

import "container/heap"

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

type MedianFinder struct {
	Low  *MaxHeap
	High *MinHeap
}

func Constructor() MedianFinder {
	low := &MaxHeap{}
	high := &MinHeap{}
	heap.Init(low)
	heap.Init(high)

	return MedianFinder{
		Low:  low,
		High: high,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// always push to low first
	heap.Push(this.Low, num)

	// Balance : lo's max must be <= hi's min
	if this.High.Len() > 0 && (*this.Low)[0] > (*this.High)[0] {
		heap.Push(this.High, heap.Pop(this.Low))
	}

	// Balance sizes: lo can have at most 1 extra element
	if this.Low.Len() > this.High.Len()+1 {
		heap.Push(this.High, heap.Pop(this.Low))
	} else if this.High.Len() > this.Low.Len() {
		heap.Push(this.Low, heap.Pop(this.High))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Low.Len() > this.High.Len() {
		return float64((*this.Low)[0])
	}
	return (float64((*this.Low)[0]) + float64((*this.High)[0])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
