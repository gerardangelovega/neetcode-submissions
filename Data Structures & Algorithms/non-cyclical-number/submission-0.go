func isHappy(n int) bool {
	m := make(map[int]struct{})
    for n != 1 {
		temp := n
		total := 0
		for temp != 0 {
			digit := temp % 10
			total = total + (digit * digit)
			temp = (temp - digit) / 10
		}
		n = total
		if _, e := m[n]; e {
			return false
		}
		m[n] = struct{}{}
	}
	return true
}
