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

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	headNext := this.head.next
	this.head.next = node
	node.prev = this.head
	headNext.prev = node
	node.next = headNext

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.lookup[key]; exists {
		node.value = value

		prev, next := node.prev, node.next
		prev.next = next
		next.prev = prev
		
		headNext := this.head.next
		this.head.next = node
		node.prev = this.head
		headNext.prev = node
		node.next = headNext
		return
	}

	if this.size == this.capacity {
		victim := this.tail.prev
		prev, next := victim.prev, victim.next
		prev.next = next
		next.prev = prev
		this.size--

		delete(this.lookup, victim.key)
	}

	node := &LRUNode{ key: key, value: value }
	this.lookup[key] = node

	headNext := this.head.next
	this.head.next = node
	node.prev = this.head
	headNext.prev = node
	node.next = headNext
	this.size++
}