package findmedianfromdatastream

// https://leetcode.cn/problems/find-median-from-data-stream/

import "container/heap"

type MedianFinder struct {
	minPQ *minHeap
	maxPQ *maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minPQ: &minHeap{},
		maxPQ: &maxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxPQ, num)
	heap.Push(this.minPQ, heap.Pop(this.maxPQ))
	if this.minPQ.Len() > this.maxPQ.Len() {
		heap.Push(this.maxPQ, heap.Pop(this.minPQ))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var res float64
	if this.maxPQ.Len() > this.minPQ.Len() {
		res = float64((*this.maxPQ)[0])
	} else {
		a := float64((*this.maxPQ)[0])
		b := float64((*this.minPQ)[0])
		res = (a + b) / 2
	}
	return res
}

type minHeap []int
type maxHeap []int

var _ heap.Interface = &minHeap{}
var _ heap.Interface = &maxHeap{}

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	mLen := m.Len()
	res := (*m)[mLen-1]
	*m = (*m)[0 : mLen-1]
	return res
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	mLen := m.Len()
	res := (*m)[mLen-1]
	*m = (*m)[0 : mLen-1]
	return res
}
