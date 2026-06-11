type Stack struct {
	arr []rune
}
func (s *Stack) Push(val rune) {
	s.arr = append(s.arr, val)
}
func (s *Stack) Pop() rune {
	n := len(s.arr)
	val := s.arr[n - 1]
	s.arr = s.arr[:n - 1]

	return val
}
func (s *Stack) Peek() rune {
	n := len(s.arr)
	val := s.arr[n - 1]

	return val
}
func (s *Stack) Length() int {
	return len(s.arr)
}

func isValid(s string) bool {
	stack := Stack{
		arr: make([]rune, 0, len(s)),
	}
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

    for _, r := range s {
		switch r {
		case ')', '}', ']':
			if stack.Length() == 0 {
				return false
			}
			if pairs[stack.Peek()] != r {
				return false
			}
			stack.Pop()
		default:
			stack.Push(r)
		}
	}

	if stack.Length() != 0 {
		return false
	}

	return true
}
