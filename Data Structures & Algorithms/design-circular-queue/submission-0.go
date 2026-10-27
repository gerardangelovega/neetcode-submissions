type CQNode struct {
	value int
	prev *CQNode
	next *CQNode
}

type MyCircularQueue struct {
	root *CQNode
	head *CQNode
	tail *CQNode
	capacity int
    size int
}


func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{
		root: nil,
		head: nil,
		tail: nil,
		capacity: k,
		size: 0,
	}
	cq.root = &CQNode{ value: -1 }
	current := cq.root
	for range k-1 {
		current.next = &CQNode{ value: -1 }
		current.next.prev = current 
		current = current.next
	}
	current.next = cq.root
	cq.root.prev = current

	cq.head = cq.root
	cq.tail = cq.root

	return cq
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.capacity { return false }

	this.tail.value = value
	this.tail = this.tail.next
	this.size++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 { return false }
	
	this.head.value = -1
	this.head = this.head.next
	this.size--
	return true
}


func (this *MyCircularQueue) Front() int {
    return this.head.value
}


func (this *MyCircularQueue) Rear() int {
	return this.tail.prev.value
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.size == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 