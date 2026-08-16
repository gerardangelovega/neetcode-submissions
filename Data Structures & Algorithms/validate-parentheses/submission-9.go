type Stack struct {
	items []rune
}
func NewStack(capacity int) Stack {
	return Stack{ items: make([]rune, 0, capacity)}
}
func (s *Stack) Push(val rune) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() rune {
	if s.Length() == 0 { return ' ' }
	res := s.items[s.Length()-1]
	s.items = s.items[:s.Length()-1]
	return res
}
func (s *Stack) Peek() rune {
	if s.Length() == 0 { return ' ' }
	res := s.items[s.Length()-1]
	return res
}
func (s *Stack) Length() int {
	return len(s.items)
}

func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := NewStack(len(s))

	for _, r := range s {
		if val, e := m[r]; e { 
			if stack.Peek() == val {
				stack.Pop()
			} else {
				return false
			}
		} else {
			stack.Push(r)
		}
	}
	if stack.Length() != 0 {
		return false
	}
	return true
}