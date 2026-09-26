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
		for {
			// If there is nothing to colide with, push to stack
			if stack.Empty() || !collide(stack.Top(), asteroid) {
				stack.Push(asteroid)
				break
			}
			// If the top asteroid is les than the current asteroid, destroy the top
			// If the top asteroid is greater than the current asteroid, destroy the current asteroid
			// If the top and current asteroid has the same size, destroy both
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

// If a has a positive direction and b has a negative direction (collision course), then return true
// Otherwise return false
func collide(a, b int) bool {
	if a > 0 && b < 0 {
		return true	
	}	
	return false
}

// If x is a negative number, then multiply it with a negative -1 to convert to positive
// Afterwards, return x
func abs(x int) int {
	if x < 0 { x = x * -1 }
	return x
}