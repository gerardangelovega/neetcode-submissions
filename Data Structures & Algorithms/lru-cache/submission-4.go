type QueueNode struct {
	key   int
	val   int
	next *QueueNode
	prev *QueueNode
}
func NewQueueNode(key, val int) *QueueNode {
	return &QueueNode{
		key:   key,
		val:   val,
		next:  nil,
		prev:  nil,
	}
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size  int
}
func NewQueue() *Queue {
	return &Queue{}
}
func (this *Queue) Enqueue(node *QueueNode) {
	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
	this.size++
}
func (this *Queue) Dequeue() *QueueNode {
	if this.size == 0 {
		return nil
	}
	node := this.head
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.head.next.prev = nil
		this.head = this.head.next
	}
	this.size--
	node.next = nil
	node.prev = nil
	return node
}
func (this *Queue) Requeue(node *QueueNode) {
	if this.size == 0 || this.size == 1 {
		return
	}
	if node.prev == nil && node.next != nil {
		// MOVE HEAD TO TAIL
		node.next.prev = nil
		this.head = node.next
	
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
		node.next = nil
		return
	}
	if node.prev != nil && node.next == nil {
		// DO NOTHING
		return	
	}
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	this.tail.next = node
	node.prev = this.tail
	this.tail = node
	node.next = nil
}

type LRUCache struct {
	lruQueue    *Queue
	lruMap 	     map[int]*QueueNode
	lruCapacity  int	
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		lruQueue:    NewQueue(),
		lruMap:      make(map[int]*QueueNode),
		lruCapacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {	
	if node, exists := this.lruMap[key]; exists {
		this.lruQueue.Requeue(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.lruMap[key]; exists {
		node.val = value
		this.lruQueue.Requeue(node)
		return
	}

	if this.lruQueue.size >= this.lruCapacity {
		node := this.lruQueue.Dequeue()
		delete(this.lruMap, node.key)
	} 
	node := NewQueueNode(key,value)
	this.lruQueue.Enqueue(node)
	this.lruMap[key] = node
}