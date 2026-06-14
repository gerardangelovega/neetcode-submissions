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
func (this *Queue) Enqueue(key, val int) *QueueNode {
	node := NewQueueNode(key, val)
	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
	this.size++
	return node
}
func (this *Queue) Dequeue() *QueueNode {
	fmt.Println(this.size)
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
	defer fmt.Println("Get Passed\n")
	fmt.Println("Starting Get")
	if this.lruMap[key] != nil {
		fmt.Println("Get Key Exists")
		node := this.lruMap[key]
		this.lruQueue.Requeue(node)
		fmt.Println("Get Key Requeued")
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	defer fmt.Println("Put Passed\n")
	if this.lruMap[key] != nil {
		fmt.Println("Put Key Updating and Requeing")
		this.lruMap[key].val = value
		this.lruQueue.Requeue(this.lruMap[key])
		fmt.Println("Put Key Updated and Requeued")
		return
	}

	if this.lruQueue.size >= this.lruCapacity {
		fmt.Println("Put Key Deleting")
		node := this.lruQueue.Dequeue()
		fmt.Println(node)
		delete(this.lruMap, node.key)
		fmt.Println("Put Key Deleted")
	} 
	fmt.Println("Put Key Inserting")
	node := this.lruQueue.Enqueue(key, value)
	this.lruMap[key] = node
	fmt.Println("Put Key Inserted")
}
