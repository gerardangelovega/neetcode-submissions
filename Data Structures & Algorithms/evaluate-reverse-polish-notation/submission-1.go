func evalRPN(tokens []string) int {
	memory := make([]int, len(tokens))
	pointer := -1

	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			pointer++
			memory[pointer] = val
		} else {
			switch token {
			case "+":
				memory[pointer - 1] = memory[pointer - 1] + memory[pointer]
				memory[pointer] = 0
				pointer--
			case "-":
				memory[pointer - 1] = memory[pointer - 1] - memory[pointer]
				memory[pointer] = 0
				pointer--
			case "*":
				memory[pointer - 1] = memory[pointer - 1] * memory[pointer]
				memory[pointer] = 0
				pointer--
			case "/":
				memory[pointer - 1] = memory[pointer - 1] / memory[pointer]
				memory[pointer] = 0
				pointer--
			default:
				fmt.Println("Error: Invalid token:", token)
			}
		}
	}

	return memory[0]
}
