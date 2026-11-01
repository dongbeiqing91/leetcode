package tryagain

import "sync"

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*LinkedNode
	head     *LinkedNode
	tail     *LinkedNode
	lock     *sync.Mutex
}

type LinkedNode struct {
	key   int
	value int
	prev  *LinkedNode
	next  *LinkedNode
}

func Constructor(capacity int) LRUCache {
	head := initLinkedNode(0, 0)
	tail := initLinkedNode(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*LinkedNode{},
		head:     head,
		tail:     tail,
		lock:     &sync.Mutex{},
	}
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) Get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if node, ok := this.cache[key]; ok {
		val := node.value
		this.moveToHead(node)
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	node := initLinkedNode(key, value)
	this.addToHead(node)
	this.cache[key] = node
	this.size++
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
		this.size--
	}
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *LinkedNode {
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
