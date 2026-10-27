type LRUNode struct {
	key int
	value int
	prev *LRUNode
	next *LRUNode
}

type LRUCache struct {
	lookup map[int]*LRUNode
    head *LRUNode
    tail *LRUNode
	capacity int
	size int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		lookup: make(map[int]*LRUNode, capacity * 2),
		head: &LRUNode{ key: -1, value: -1 },
		tail: &LRUNode{ key: -1, value: -1 },
		capacity: capacity,
		size: 0,
	} 
	lru.head.next = lru.tail
	lru.head.prev = lru.tail

	lru.tail.next = lru.head
	lru.tail.prev = lru.head

	return lru
}

func (this *LRUCache) Get(key int) int {
    node, exists := this.lookup[key]
	if !exists { 
		return -1 
	}
	this.requeue(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.lookup[key]; exists {
		node.value = value
		this.requeue(node)
		return
	}

	if this.size == this.capacity {
		victim := this.dequeue()
		delete(this.lookup, victim.key)
		this.size--
	}

	node := &LRUNode{ key: key, value: value }
	this.lookup[key] = node
	this.enqueue(node)
	this.size++
}

func (this *LRUCache) detach(node *LRUNode) *LRUNode {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	return node
}

func (this *LRUCache) enqueue(node *LRUNode) {
	prev, next := this.head, this.head.next

	prev.next = node
	node.prev = prev

	next.prev = node
	node.next = next
}

func (this *LRUCache) dequeue() *LRUNode {
	node := this.tail.prev

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	return node
}

func (this *LRUCache) requeue(node *LRUNode) {
	node = this.detach(node)
	this.enqueue(node)
}