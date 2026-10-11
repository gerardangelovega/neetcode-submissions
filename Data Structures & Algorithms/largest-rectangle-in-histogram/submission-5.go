type Stack struct {
	items []int	
}
func NewStack() Stack {
	return Stack{
		items: []int{},
	}
}
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() int {
	if s.Length() == 0 { return -1 }
	num := s.items[s.Length()-1]
	s.items = s.items[:s.Length()-1]
	return num
}
func (s *Stack) Length() int {
	return len(s.items)
}
func (s *Stack) Top() int {
	if s.Length() == 0 { return -1 }
	return s.items[s.Length()-1]
}

func largestRectangleArea(heights []int) int {
	n := len(heights)

	left, right := NewStack(), NewStack()
	leftmost, rightmost := make([]int, n), make([]int, n)
	for i := range n { leftmost[i], rightmost[i] = -1, n }

	for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
		for left.Length() > 0 && heights[left.Top()] >= heights[l] { left.Pop() }
		if left.Length() > 0 { leftmost[l] = left.Top() }
		left.Push(l)

		for right.Length() > 0 && heights[right.Top()] >= heights[r] { right.Pop() }
		if right.Length() > 0 { rightmost[r] = right.Top() }
		right.Push(r)
	}

	res := 0
	for i := range n {
		leftmost[i]++
		rightmost[i]--
		res = max(res, heights[i] * (rightmost[i] - leftmost[i] + 1))
	}

	return res
}

// L2R: 1,2,2,4
// R2L: 1,7

// L2R: 1,3,7
// R2L: 1

// L2R: 1,2,3
// R2L: 1,2