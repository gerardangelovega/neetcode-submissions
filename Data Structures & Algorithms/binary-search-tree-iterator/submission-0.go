/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack struct {
	items []*TreeNode
}
func NewStack() Stack {
	return Stack{
		items: make([]*TreeNode, 0),
	}
}
func (s *Stack) Push(val *TreeNode) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() { return nil }
	val := s.items[s.Length()-1]
	s.items = s.items[:s.Length()-1]
	return val
}
func (s *Stack) Peek() *TreeNode {
	if s.IsEmpty() { return nil }
	return s.items[s.Length()-1]
}
func (s *Stack) Length() int {
	return len(s.items)
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

type BSTIterator struct {
	stack Stack	
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator {
		stack: NewStack(),
	}
	curr := root
	for curr != nil {
		it.stack.Push(curr)
		curr = curr.Left
	}
	return it
}

func (this *BSTIterator) Next() int {
	top := this.stack.Pop()
	if top.Right != nil {
		curr := top.Right
		for curr != nil {
			this.stack.Push(curr)
			curr = curr.Left
		}
	}
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.IsEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 