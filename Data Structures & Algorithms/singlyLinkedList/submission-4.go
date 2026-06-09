type Node struct {
    next* Node
    val int
}

func NewNode(val int) *Node {
    return &Node{
        next: nil,
        val: val,
    }
}

type LinkedList struct {
    head* Node
    tail* Node
    len int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: nil,
        tail: nil,
        len: 0,
    }
}

func (ll *LinkedList) Get(index int) int {
    i := 0

    for curr := ll.head; curr != nil; curr, i = curr.next, i + 1 {
        if i == index {
            return curr.val
        }
    }

    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    node := NewNode(val)

    if ll.head == nil && ll.tail == nil {
        ll.head = node
        ll.tail = node
        ll.len = ll.len + 1
        return
    }

    node.next = ll.head
    ll.head = node

    ll.len = ll.len + 1
}

func (ll *LinkedList) InsertTail(val int) {
    node := NewNode(val)

    if ll.head == nil && ll.tail == nil {
        ll.head = node
        ll.tail = node
        ll.len = ll.len + 1
        return
    }

    ll.tail.next = node
    ll.tail = node

    ll.len = ll.len + 1
}

func (ll *LinkedList) Remove(index int) bool {
    curr := ll.head
    prev := curr

    for i := 0; i != index; i, prev, curr = i + 1, curr, curr.next {
        if curr == nil {
            return false
        }
    }

    if curr == nil {
        return false
    }

    if index == 0 {
        ll.head = ll.head.next
        ll.len = ll.len - 1
        return true
    }

    if (index == ll.len - 1) {
        prev.next = nil
        ll.tail = prev
        ll.len = ll.len - 1

        return true
    }

    prev.next = curr.next
    ll.len = ll.len - 1

    return true
}

func (ll *LinkedList) GetValues() []int {
    arr := make([]int, ll.len)

    i := 0

    for curr := ll.head; curr != nil; curr, i = curr.next, i + 1 {
        arr[i] = curr.val
    }

    return arr
}
