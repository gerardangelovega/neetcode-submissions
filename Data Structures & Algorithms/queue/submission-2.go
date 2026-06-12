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
	d.appendFunc(value, func(node *DequeNode) {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	})
}
func (d *Deque) AppendLeft(value int) {
	d.appendFunc(value, func(node *DequeNode) {
		node.next = d.head
		d.head.prev = node
		d.head = node
	})
}
func (d *Deque) Pop() int {
	return d.popFunc(func() int {
		node := d.tail.prev
		prev := d.tail
		prev.prev = nil
		node.next = nil
		d.tail = node
		return prev.val
	})
}
func (d *Deque) PopLeft() int {
	return d.popFunc(func() int {
		node := d.head.next
		prev := d.head
		prev.next = nil
		node.prev = nil
		d.head = node
		return prev.val
	})
}
func (d *Deque) appendFunc(value int, fn func(*DequeNode)) {
	node := NewNode(value)
	if d.size == 0 {
		d.head = node
		d.tail = node
	} else {
		fn(node)
	}
	d.size++
}
func (d *Deque) popFunc(fn func() int) int {
	if d.size == 0 {
		return -1
	}
	val := 0
	if d.size == 1 {
		val = d.head.val;
		d.head = nil
		d.tail = nil
	} else {
		val = fn()
	}
	d.size--
	return val
}