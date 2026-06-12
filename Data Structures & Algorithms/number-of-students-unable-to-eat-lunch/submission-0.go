type Node struct {
	val   int
	next *Node
}
func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type Queue struct {
	head *Node
	tail *Node
	size  int
}
func NewQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}
func (q *Queue) Enqueue(val int) {
	node := NewNode(val)
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}
func (q *Queue) Dequeue() int {
	if q.size == 0 {
		return -1
	}	
	node := q.head
	if q.size == 1 {
		q.head = nil
		q.tail = nil	
	} else {
		q.head = q.head.next
	}
	q.size--
	return node.val
}
func (q *Queue) Requeue() {
	if q.size == 1 || q.size == 0 {
		return
	}
	node := q.head

	q.head = q.head.next
	q.tail.next = node
	q.tail = node
	q.tail.next = nil
}
func (q *Queue) Size() int {
	return q.size
}
func (q *Queue) Peek() int {
	return q.head.val
}

func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	student_queue  := NewQueue()
	sandwich_queue := NewQueue()

	for i := 0; i < n; i++ {
		student_queue.Enqueue(students[i])
		sandwich_queue.Enqueue(sandwiches[i])
	}

	student_requeues := 0
	for true {
		// loop termination logic
		if student_queue.Size() == 0 {
			break
		}
		if sandwich_queue.Size() == 0 {
			break
		}
		if student_requeues > student_queue.Size() {
			break
		}

		// algorithm logic
		if student_queue.Peek() == sandwich_queue.Peek() {
			student_queue.Dequeue()
			sandwich_queue.Dequeue()
			student_requeues = 0
		} else {
			student_queue.Requeue()
			student_requeues++
		}
	}
	
	return student_queue.Size()
}