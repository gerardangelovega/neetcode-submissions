type MyStackNode struct {
	val  int
	next *MyStackNode
}
func NewNode(val int) *MyStackNode {
	return &MyStackNode{
		val: val,
	}
}

type MyStack struct {
	head *MyStackNode
	tail *MyStackNode
	size  int
}
func Constructor() MyStack {
	return MyStack{
		head: nil,
		tail: nil,
		size: 0,
	}
}
func (this *MyStack) Push(x int) {
	node := NewNode(x)
	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.size++

	this.rotate(this.size - 1)
}
func (this *MyStack) Pop() int {
	if this.size == 0 {
		return -1
	}
	val := this.head.val
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
	}
	this.size--	
	return val
}
func (this *MyStack) Top() int {
	return this.head.val
}
func (this *MyStack) Empty() bool {
	return this.size == 0
}
func (this *MyStack) rotate(steps int) {
	if this.size == 0 || this.size == 1 {
		return
	}

	for i := 0; i < steps; i++ {
		node := this.head

		this.head = this.head.next
		this.tail.next = node
		this.tail = node
		this.tail.next = nil
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
