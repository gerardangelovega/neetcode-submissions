func isPerfectSquare(num int) bool {
	l, r, m := 0, num, num / 2

	for l <= r {
		m = l + (r - l)	/ 2
		msq := m * m
		
		if msq == num {
			return true
		} else if msq < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
