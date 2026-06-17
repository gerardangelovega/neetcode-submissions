func confusingNumber(n int) bool {
	m := make(map[int]struct{})

	if n == 0 {
		m[n] = struct{}{}
	}

	for n != 0 {
		num := n % 10

		switch num {
		case 0, 1, 6, 7, 8, 9:
			m[num] = struct{}{}
		default:
			return false
		}

		n = n / 10
	}

	if _, exists := m[0]; exists && len(m) == 1 {
		return false
	}
	if _, exists := m[1]; exists && len(m) == 1 {
		return false
	}

	return true
}
