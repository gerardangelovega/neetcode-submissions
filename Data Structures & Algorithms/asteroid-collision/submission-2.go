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
	if len(s.items) == 0 { return -1 }
	rv := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return rv
}
func (s *Stack) Top() int {
	if len(s.items) == 0 { return -1 }
	return s.items[len(s.items)-1]
}
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func asteroidCollision(asteroids []int) []int {
	stack := NewStack()
	for _, asteroid := range asteroids {
		if stack.Empty() {
			stack.Push(asteroid)
			continue
		}
		if !collide(stack.Top(), asteroid) {
			stack.Push(asteroid)
			continue
		}
		for {
			if stack.Empty() || !collide(stack.Top(), asteroid) {
				stack.Push(asteroid)
				break
			}
			if abs(stack.Top()) < abs(asteroid) {
				stack.Pop()
			} else if abs(stack.Top()) == abs(asteroid) {
				stack.Pop()
				break
			} else {
				break
			}
		}
	}
	return stack.items
}

func collide(a, b int) bool {
	if a > 0 && b < 0 {
		return true	
	}	
	return false
}

func abs(x int) int {
	if x < 0 { x = x * -1 }
	return x
}