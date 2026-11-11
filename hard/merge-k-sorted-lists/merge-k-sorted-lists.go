package mergeksortedlists

// https://leetcode.cn/problems/merge-k-sorted-lists/

import "container/heap"

var _ heap.Interface = &PriorityQueue{}

type PriorityQueue []*ListNode

// Len implements [heap.Interface].
func (p PriorityQueue) Len() int {
	return len(p)
}

// Less implements [heap.Interface].
func (p PriorityQueue) Less(i int, j int) bool {
	return p[i].Val < p[j].Val
}

// Pop implements [heap.Interface].
func (p *PriorityQueue) Pop() any {
	pLen := p.Len()
	res := (*p)[pLen-1]
	*p = (*p)[:pLen-1]
	return res
}

// Push implements [heap.Interface].
func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(*ListNode))
}

// Swap implements [heap.Interface].
func (p *PriorityQueue) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, item := range lists {
		if item != nil {
			heap.Push(pq, item)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for pq.Len() > 0 {
		n := heap.Pop(pq).(*ListNode)
		tail.Next = &ListNode{
			Val: n.Val,
		}
		tail = tail.Next
		if n.Next != nil {
			heap.Push(pq, n.Next)
		}
	}
	return dummy.Next
}
