func isValid(s string) bool {
    stack := []rune{}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)	
		case ')', '}', ']':
			if (len(stack) == 0) {
				return false
			}
			if stack[len(stack) - 1] == '(' && r == ')' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack) - 1] == '[' && r == ']' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack) - 1] == '{' && r == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}
