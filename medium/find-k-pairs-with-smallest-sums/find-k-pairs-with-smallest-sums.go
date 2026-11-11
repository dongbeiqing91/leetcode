package findkpairswithsmallestsums

import "container/heap"

// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/

var _ heap.Interface = &PriorityQueue{}

type PriorityQueue [][]int

// Len implements [heap.Interface].
func (p PriorityQueue) Len() int {
	return len(p)
}

// Less implements [heap.Interface].
func (p PriorityQueue) Less(i int, j int) bool {
	return p[i][0]+p[i][1] < p[j][0]+p[j][1]
}

// Pop implements [heap.Interface].
func (p *PriorityQueue) Pop() any {
	pLen := len(*p)
	res := (*p)[pLen-1]
	*p = (*p)[0 : pLen-1]
	return res
}

// Push implements [heap.Interface].
func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.([]int))
}

// Swap implements [heap.Interface].
func (p *PriorityQueue) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}
	var res [][]int
	for len(*pq) > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		k--
		nextIndex := cur[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})
		}
		n := []int{cur[0], cur[1]}
		res = append(res, n)
	}
	return res
}
