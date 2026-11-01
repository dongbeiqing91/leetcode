package medium

type LRUCache struct {
    size, capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
    head := initDLinkedNode(0, 0)
	tail := initDLinkedNode(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		size: 0,
		capacity: capacity,
		cache: map[int]*DLinkedNode{},
		head: head,
		tail: tail,
	}
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		val := node.value
		this.moveToHead(node)
		return val
	}
}

func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	node := initDLinkedNode(key, value)
	this.addToHead(node)
	this.cache[key] = node
	this.size ++
	if this.size > this.capacity {
		removedNode := this.removeTail()
		delete(this.cache, removedNode.key)
		this.size --
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	prevNode := node.prev
	nextNode := node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (this *LRUCache) removeTail() *DLinkedNode {
	last := this.tail.prev
	this.removeNode(last)
	return last
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */