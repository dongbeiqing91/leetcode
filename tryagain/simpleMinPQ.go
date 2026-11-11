package tryagain

type SimpleMinPQ struct {
	heap []int
	size int
}

func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

func (pq *SimpleMinPQ) left(node int) int {
	return node*2 + 1
}

func (pq *SimpleMinPQ) right(node int) int {
	return node*2 + 2
}

func (pq *SimpleMinPQ) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *SimpleMinPQ) peek() int {
	return pq.heap[0]
}

func (pq *SimpleMinPQ) push(x int) {
	pq.heap[pq.size] = x
	pq.swim(pq.size)
	pq.size++
}

func (pq *SimpleMinPQ) pop() int {
	res := pq.heap[0]
	pq.heap[0] = pq.heap[pq.size-1]
	pq.size--
	pq.sink(0)
	return res
}

func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.heap[pq.parent(node)] > pq.heap[node] {
		pq.swap(pq.parent(node), node)
		node = pq.parent(node)
	}
}

func (pq *SimpleMinPQ) sink(node int) {
	for pq.left(node) < pq.size || pq.right(node) < pq.size {
		min := node
		if pq.left(node) < pq.size && pq.heap[min] > pq.heap[pq.left(node)] {
			min = pq.left(node)
		}
		if pq.right(node) < pq.size && pq.heap[min] > pq.heap[pq.right(node)] {
			min = pq.right(node)
		}
		if min == node {
			break
		}
		pq.swap(min, node)
		node = min
	}
}
