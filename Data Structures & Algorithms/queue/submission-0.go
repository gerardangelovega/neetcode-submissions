type DequeNode struct {
	val   int
	next *DequeNode
	prev *DequeNode
}
func NewNode(val int) *DequeNode {
	return &DequeNode{
		val: val,
	}
}

type Deque struct {
	head *DequeNode
	tail *DequeNode
	size  int
}
func NewDeque() *Deque {
	return &Deque{}
}
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}
func (d *Deque) Append(value int) {
	node := NewNode(value)
	if d.size == 0 {
		d.appendEmpty(node)
	} else {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.size++
}
func (d *Deque) AppendLeft(value int) {
	node := NewNode(value)
	if d.size == 0 {
		d.appendEmpty(node)
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.size++
}
func (d *Deque) Pop() int {
	if d.size == 0 {
		return -1
	}
	val := d.tail.val
	if d.size == 1 {
		d.deleteLast()
	} else {
		node := d.tail.prev
		d.tail.prev = nil
		node.next = nil
		d.tail = node
	}
	d.size--
	return val
}
func (d *Deque) PopLeft() int {
	if d.size == 0 {
		return -1
	}
	val := d.head.val
	if d.size == 1 {
		d.deleteLast()
	} else {
		node := d.head.next
		d.head.next = nil
		node.prev = nil
		d.head = node
	}
	d.size--
	return val
}
func (d *Deque) appendEmpty(node *DequeNode) {
	d.head = node
	d.tail = node
}
func (d *Deque) deleteLast() {
	d.head = nil
	d.tail = nil
}