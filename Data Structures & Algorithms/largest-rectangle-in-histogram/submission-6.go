type Stack struct {
	items []int
}
func NewStack(capacity int) Stack {
	return Stack {
		items: make([]int, 0, capacity),
	}
}
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() int {
	if s.Len() == 0 { return -1 }
	val := s.Top()
	s.items = s.items[:len(s.items)-1]
	return val
}
func (s *Stack) Top() int {
	if s.Len() == 0 { return -1 }
	return s.items[len(s.items)-1]
}
func (s *Stack) Len() int {
	return len(s.items)
}

func largestRectangleArea(heights []int) int {
	n := len(heights)

	pse, nse := make([]int, n), make([]int, n)
	for i := range n {
		pse[i], nse[i] = -1, n
	}

	left, right := NewStack(n), NewStack(n)
	for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
		for left.Len() > 0 && heights[left.Top()] >= heights[l] {
			left.Pop()
		}
		if left.Len() > 0 {
			pse[l] = left.Top()
		}
		left.Push(l)

		for right.Len() > 0 && heights[right.Top()] >= heights[r] {
			right.Pop()
		}
		if right.Len() > 0 {
			nse[r] = right.Top()
		}
		right.Push(r)
	}

	res := 0
	for i := range n {
		pse[i]++
		nse[i]--
		res = max(res, heights[i] * (nse[i] - pse[i] + 1))
	}
	return res
}
