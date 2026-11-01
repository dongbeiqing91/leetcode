package medium

import "sync"

type ThreadSafeLRUNode struct {
	key, value int
	prev, next *ThreadSafeLRUNode
}

type ThreadSafeLRUCache struct {
	mu       sync.Mutex
	size     int
	capacity int
	cache    map[int]*ThreadSafeLRUNode
	head     *ThreadSafeLRUNode
	tail     *ThreadSafeLRUNode
}

func NewThreadSafeLRUCache(capacity int) *ThreadSafeLRUCache {
	if capacity < 0 {
		capacity = 0
	}

	head := &ThreadSafeLRUNode{}
	tail := &ThreadSafeLRUNode{}
	head.next = tail
	tail.prev = head

	return &ThreadSafeLRUCache{
		capacity: capacity,
		cache:    make(map[int]*ThreadSafeLRUNode),
		head:     head,
		tail:     tail,
	}
}

func (c *ThreadSafeLRUCache) Get(key int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.moveToHead(node)
	return node.value
}

func (c *ThreadSafeLRUCache) Put(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity == 0 {
		return
	}

	if node, ok := c.cache[key]; ok {
		node.value = value
		c.moveToHead(node)
		return
	}

	node := &ThreadSafeLRUNode{key: key, value: value}
	c.addToHead(node)
	c.cache[key] = node
	c.size++

	if c.size > c.capacity {
		removed := c.removeTail()
		delete(c.cache, removed.key)
		c.size--
	}
}

func (c *ThreadSafeLRUCache) moveToHead(node *ThreadSafeLRUNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *ThreadSafeLRUCache) addToHead(node *ThreadSafeLRUNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *ThreadSafeLRUCache) removeNode(node *ThreadSafeLRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *ThreadSafeLRUCache) removeTail() *ThreadSafeLRUNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
