package kthsmallestelementinasortedmatrix

import "container/heap"

var _ heap.Interface = &PriorityQueue{}

type PriorityQueue [][]int

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.([]int))
}

func (p *PriorityQueue) Pop() any {
	pLen := p.Len()
	res := (*p)[pLen-1]
	*p = (*p)[0 : pLen-1]
	return res
}

func kthSmallest(matrix [][]int, k int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	maxCol := len(matrix)
	for i, item := range matrix {
		heap.Push(pq, []int{item[0], i, 0})
	}
	var res int
	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		if k == 1 {
			res = cur[0]
			break
		}
		row := cur[1]
		col := cur[2]
		if col+1 < maxCol {
			next := matrix[row][col+1]
			heap.Push(pq, []int{next, row, col + 1})
		}
		k--
	}
	return res
}
