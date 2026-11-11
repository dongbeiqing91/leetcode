package priorityqueue

// https://leetcode.cn/problems/kth-largest-element-in-a-stream/

import (
	"container/heap"
)

var _ heap.Interface = &PriorityQueue{}

type PriorityQueue []int

// Len implements [heap.Interface].
func (p *PriorityQueue) Len() int {
	return len(*p)
}

// Less implements [heap.Interface].
func (p *PriorityQueue) Less(i int, j int) bool {
	return (*p)[i] < (*p)[j]
}

// Pop implements [heap.Interface].
func (p *PriorityQueue) Pop() any {
	last := (*p)[p.Len() - 1]
	*p = (*p)[:p.Len() - 1]
	return last
}

// Push implements [heap.Interface].
func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(int))
}

// Swap implements [heap.Interface].
func (p *PriorityQueue) Swap(i int, j int) {
	q := *p
	q[i], q[j] = q[j], q[i]
}

type KthLargest struct {
	k  int
	pq *PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, n := range nums {
		heap.Push(pq, n)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	return KthLargest{
		k:  k,
		pq: pq,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.pq, val)
	if this.pq.Len() > this.k {
		heap.Pop(this.pq)
	}
	return (*this.pq)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
