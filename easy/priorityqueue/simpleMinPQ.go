package priorityqueue

// https://labuladong.online/zh/algo/data-structure-basic/binary-heap-implement/

type SimpleMinPQ struct {
	// 底层使用数组实现二叉堆
	heap []int

	// 堆中元素的数量
	size int
}

// 父节点的索引
func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

// 左子节点的索引
func (pq *SimpleMinPQ) left(node int) int {
	return node*2 + 1
}

// 右子节点的索引
func (pq *SimpleMinPQ) right(node int) int {
	return node*2 + 2
}

// 交换数组的两个元素
func (pq *SimpleMinPQ) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

// 查，返回堆顶元素，时间复杂度 O(1)
func (pq *SimpleMinPQ) peek() int {
	return pq.heap[0]
}

// 增，向堆中插入一个元素，时间复杂度 O(logN)
func (pq *SimpleMinPQ) push(x int) {
	// 把新元素追加到最后
	pq.heap[pq.size] = x
	// 然后上浮到正确位置
	pq.swim(pq.size)
	pq.size++
}

// 删，删除堆顶元素，时间复杂度 O(logN)
func (pq *SimpleMinPQ) pop() int {
	res := pq.heap[0]
	// 把堆底元素放到堆顶
	pq.heap[0] = pq.heap[pq.size-1]
	pq.size--
	// 然后下沉到正确位置
	pq.sink(0)
	return res
}

// 上浮操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.heap[pq.parent(node)] > pq.heap[node] {
		pq.swap(pq.parent(node), node)
		node = pq.parent(node)
	}
}

// sink 的前提是下标 node 处的左右子树已经各自满足小根堆规则。
// 如果有孩子比当前元素小，就和最小的孩子交换。交换后，只有被换下去的元素
// 还可能比它的新孩子大，因此只沿它所在的分支继续检查。单次下沉最多走一个树高，
// 时间复杂度为 O(log N)。
//
// pop 从合法的堆中把最后一个叶子的值移到堆顶。去掉原来的叶子位置后，
// 左右子树仍然各自成堆，所以调用一次 sink(0) 就能修复整个堆。
//
// 从无序数组建堆时，要从最后一个非叶子节点（下标 pq.size/2 - 1）开始，
// 按下标递减到 0，对每个非叶子节点调用一次 sink。叶子没有孩子，天然满足堆规则。
// 孩子的下标比父节点大，这个顺序保证处理父节点时，它的左右子树已经成堆。
func (pq *SimpleMinPQ) sink(node int) {
	// 孩子的下标小于 size 才有效。只要有一个孩子，就继续检查。
	for pq.left(node) < pq.size || pq.right(node) < pq.size {
		// 先把当前位置作为候选。min 保存的是下标，pq.heap[min] 才是值。
		min := node
		// 先确认左孩子存在，再比较值。左孩子不存在时，&& 会跳过右侧的比较。
		if pq.left(node) < pq.size && pq.heap[pq.left(node)] < pq.heap[min] {
			// 左孩子更小，把候选下标更新为左孩子的下标。
			min = pq.left(node)
			// 左孩子检查完毕，min 指向当前节点和有效左孩子中较小的那个。
		}
		// 右孩子存在时，将它的值和当前候选的值比较，候选也可能是左孩子。
		if pq.right(node) < pq.size && pq.heap[pq.right(node)] < pq.heap[min] {
			// 右孩子比当前候选更小，把候选下标更新为右孩子的下标。
			min = pq.right(node)
			// 右孩子检查完毕，min 保存这一轮最小元素的下标。
		}
		// 如果候选仍是当前位置，说明两个孩子中没有值比当前元素更小。
		if min == node {
			// 当前元素不大于任何孩子，且左右子树已各自成堆，整棵子树就满足堆规则。
			// 如果子树原本有问题，这里仍会退出。例如对 [1, 4, 3, 2] 调用 sink(0)，
			// 根节点 1 比两个孩子都小，但更深一层的 4 > 2 不会被修复。
			break
			// 提前退出分支到此结束。
		}
		// 让值更小的孩子成为父节点，使它同时不大于左右孩子。
		pq.swap(node, min)
		// 原来的元素已经移到 min，下一轮从这个下标继续检查。
		node = min
		// 回到循环开头，检查元素的新位置是否还有孩子。
	}
	// 到达叶子或当前位置满足堆规则时，下沉结束。
}
