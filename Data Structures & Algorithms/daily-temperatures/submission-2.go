func pop(arr *[]int) int {
	item := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return item
}

func push(arr *[]int, val int) {
	*arr = append(*arr, val)
}

func peek(arr *[]int) int {
	return (*arr)[len(*arr)-1]
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[peek(&stack)] {
			index := pop(&stack)
			res[index] = i - index
		}
		push(&stack, i)
	}

	return res
}
