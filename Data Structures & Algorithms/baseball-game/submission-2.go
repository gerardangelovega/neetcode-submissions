type Stack struct {
	arr []int	
	head int
}
func (s *Stack) Push(val int) {
	s.head++
	if s.head >= len(s.arr) {
		s.arr = append(s.arr, val)
	} else {
		s.arr[s.head] = val
	}
}
func (s *Stack) Pop() int {
	if s.head < 0 {
		return -1
	}

	val := s.arr[s.head]
	s.arr[s.head] = 0
	s.head--

	return val
}
func (s *Stack) Peek() int {
	val := s.arr[s.head]

	return val
}
func (s *Stack) Length() int {
	return s.head + 1
}
func (s *Stack) Sum2() int {
	return s.arr[s.head] + s.arr[s.head - 1]
}
func (s *Stack) SumAll() int {
	total := 0

	for i := range s.Length() {
		total = total + s.arr[i]
	}

	return total
}
func NewStack() *Stack {
	return &Stack {
		arr: make([]int, 0, 4),
		head: -1,
	}
}

func calPoints(operations []string) int {
	stack := NewStack()

	for _, s := range operations {
		i, err := strconv.Atoi(s)
		if err == nil {
			stack.Push(i)
			continue
		}
		switch s {
		case "+":
			stack.Push(stack.Sum2())
		case "C":
			stack.Pop()
		case "D":
			stack.Push(stack.Peek() * 2)
		}
	}

	return stack.SumAll()
}
