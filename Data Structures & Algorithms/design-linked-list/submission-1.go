type MyLinkedList struct {
	head  *Node
	tail  *Node
	length int
}

type Node struct {
	val   int
	next *Node
	prev *Node
}


func Constructor() MyLinkedList {
    return MyLinkedList{
		head: nil,
		tail: nil,
		length: 0,
	}
}


func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}

	curr := this.head

	for i := 0; i <= index; i++ {
		if curr == nil {
			break
		}
		if i == index {
			return curr.val
		}
		curr = curr.next
	}

	return -1
}

func (this *MyLinkedList) AddEmpty(val int)  {
    node := &Node{
		val: val,
		next: nil,
		prev: nil,
	}
	this.head = node
	this.tail = node
	this.length++
}

func (this *MyLinkedList) AddAtHead(val int)  {
	if this.length == 0 {
		this.AddEmpty(val)
		return	
	}

    node := &Node{
		val: val,
		next: this.head,
		prev: nil,
	}
	this.head.prev = node
	this.head = node
	this.length++
}


func (this *MyLinkedList) AddAtTail(val int)  {
	if this.length == 0 {
		this.AddEmpty(val)
		return	
	}

    node := &Node{
		val: val,
		next: nil,
		prev: this.tail,
	}
	this.tail.next = node
	this.tail = node
	this.length++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.length {
		return
	}
	if this.length == 0 {
		this.AddEmpty(val)
		return	
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}

	curr := this.head
	for i := 0; i <= index; i++ {
		if curr == nil {
			break
		}
		if i == index {
			prev := curr.prev
			next := curr
			node := &Node {
				val: val,
				next: next,
				prev: prev,
			}
			prev.next = node
			next.prev = node
			this.length++
			break
		}
		curr = curr.next
	}
}

func (this *MyLinkedList) DeleteLast()  {
	this.head = nil
	this.tail = nil
	this.length--
}

func (this *MyLinkedList) DeleteAtHead()  {
	this.head.next.prev = nil	
	this.head = this.head.next
	this.length--
}

func (this *MyLinkedList) DeleteAtTail()  {
	this.tail.prev.next = nil
	this.tail = this.tail.prev
	this.length--
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	fmt.Println(this.length)
	if index >= this.length || this.length == 0 {
		return
	}
	if this.length == 1 {
		this.DeleteLast()
		return
	}
	if index == 0 {
		this.DeleteAtHead()
		return
	}
	if index == this.length - 1 {
		this.DeleteAtTail()
		return
	}

	curr := this.head
	for i := 0; i <= index; i++ {
		if curr == nil {
			break
		}
		if i == index {
			prev := curr.prev
			next := curr.next

			prev.next = next
			next.prev = prev
			this.length--
			fmt.Println(prev.val, next.val)
			break
		}
		curr = curr.next
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */