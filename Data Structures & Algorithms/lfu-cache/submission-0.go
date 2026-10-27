type LFUNode struct {
	key int
	value int
	frequency int

	prev *LFUNode
	next *LFUNode
}

type LFUList struct {
	head *LFUNode
	tail *LFUNode
	size int
}
func NewList() *LFUList {
	list := &LFUList{
		head: &LFUNode{ key: -1, value: -1, frequency: -1 },
		tail: &LFUNode{ key: -1, value: -1, frequency: -1 },
		size: 0,
	}

	list.head.prev = list.tail
	list.head.next = list.tail

	list.tail.prev = list.head
	list.tail.next = list.head

	return list
}
func (l *LFUList) detach(node *LFUNode) *LFUNode {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	l.size--

	return node
}
func (l *LFUList) enqueue(node *LFUNode) {
	prev, next := l.head, l.head.next

	prev.next = node
	node.prev = prev

	next.prev = node
	node.next = next

	l.size++
}
func (l *LFUList) dequeue() *LFUNode {
	node := l.tail.prev
	return l.detach(node)
}

type LFUCache struct {
    nMap map[int]*LFUNode
    fMap map[int]*LFUList

	leastFrequency int
	mostFrequency int

	capacity int
	size int
}
func Constructor(capacity int) LFUCache {
	return LFUCache{
		nMap: make(map[int]*LFUNode),
		fMap: make(map[int]*LFUList),
		leastFrequency: math.MaxInt,
		mostFrequency: math.MinInt,
		capacity: capacity,
		size: 0,
	}
}
func (this *LFUCache) Get(key int) int {
	node, exists := this.nMap[key]
    if !exists {
		return -1
	}

	this.Update(node)

	return node.value
}

func (this *LFUCache) Put(key int, value int)  {
    if node, exists := this.nMap[key]; exists {
		node.value = value;
		this.Update(node)
		return
	}

	if this.size == this.capacity {
		listVictim, exists := this.fMap[this.leastFrequency]
		for (!exists || listVictim.size == 0) && this.leastFrequency <= this.mostFrequency {
			this.leastFrequency++
			listVictim, exists = this.fMap[this.leastFrequency]
		}
		victim := listVictim.dequeue()
		delete(this.nMap, victim.key)
		this.size--
	}

	node := &LFUNode{
		key: key,	
		value: value,	
		frequency: 1,
	}

	list, exists := this.fMap[node.frequency]
	if !exists {
		this.fMap[node.frequency] = NewList()
		list = this.fMap[node.frequency]
	}
	list.enqueue(node)
	this.nMap[node.key] = node
	this.leastFrequency = min(this.leastFrequency, node.frequency)
	this.mostFrequency = max(this.mostFrequency, node.frequency)
	this.size++
}

func (this *LFUCache) Update(node* LFUNode) {
	listOld, exists := this.fMap[node.frequency]
	if exists {
		listOld.detach(node)
	}

	node.frequency++
	listNew, exists := this.fMap[node.frequency]
	if !exists {
		this.fMap[node.frequency] = NewList()
		listNew = this.fMap[node.frequency]
	}
	listNew.enqueue(node)

	this.leastFrequency = min(this.leastFrequency, node.frequency)
	this.mostFrequency = max(this.mostFrequency, node.frequency)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
