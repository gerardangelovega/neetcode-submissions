type Stack struct {
	arr []int	
}
func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
}
func (s *Stack) Pop() int {
	n := len(s.arr)
	val := s.arr[n-1]
	s.arr = s.arr[:n-1]

	return val
}
func (s *Stack) Peek() int {
	n := len(s.arr)
	val := s.arr[n-1]

	return val
}
func (s *Stack) Length() int {
	return len(s.arr)
}
func (s *Stack) Sum2() int {
	n := len(s.arr)
	return s.arr[n-1] + s.arr[n-2]
}
func (s *Stack) SumAll() int {
	total := 0

	for i := range len(s.arr) {
		total = total + s.arr[i]
	}

	return total
}

func calPoints(operations []string) int {
	stack := Stack{
		arr: make([]int, 0, len(operations)),
	}

	for _, s := range operations {
		switch s {
		case "+":
			stack.Push(stack.Sum2())
		case "C":
			stack.Pop()
		case "D":
			stack.Push(stack.Peek() * 2)
		default:
			i, _ := strconv.Atoi(s)
			stack.Push(i)
		}
	}

	return stack.SumAll()
}