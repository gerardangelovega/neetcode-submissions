type Stack struct {
	items []int
}
func NewStack() Stack {
	return Stack{
		items: make([]int, 0, 4),
	}
}
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	rv := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return rv
}
func (s *Stack) Top() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

type MyQueue struct {
	push Stack
	pop Stack
}

func Constructor() MyQueue {
	return MyQueue{
		push: NewStack(),
		pop: NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.push.Push(x);
}

func (this *MyQueue) Pop() int {
	if this.pop.Empty() {
		for !this.push.Empty() {
			this.pop.Push(this.push.Pop())
		}
	}
	return this.pop.Pop()
}

func (this *MyQueue) Peek() int {
	if this.pop.Empty() {
		for !this.push.Empty() {
			this.pop.Push(this.push.Pop())
		}
	}
	return this.pop.Top()
}

func (this *MyQueue) Empty() bool {
	return this.push.Empty() && this.pop.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */

// stack 1, stack 2
// Push 3 times
// 1, 
// 1, 2
// 1, 2, 3
// Pop
// 1, 2,	3
// 1,		3, 2
// 			3, 2, 1(out)
//			3, 2
// 2		3
// 2, 3		_