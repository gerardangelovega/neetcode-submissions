func plusOne(digits []int) []int {
	n := len(digits)
	c := 0
	for i := n-1; i >= 0; i-- {
		if i == n-1 {
			c = (digits[i] + 1) / 10
			digits[i] = (digits[i] + 1) % 10
			continue
		}
		if c == 0 {
			break
		}
		temp := digits[i] + c
		digits[i] = temp % 10
		c = temp / 10
	}
	if c != 0 {
		res := make([]int, n + 1)
		res[0] = c
		copy(res[1:], digits)
		return res
	}
	return digits
}
