func mySqrt(x int) int {
	res := 0

	l, r := 0, x
	for l <= r {
		m := (l + r) / 2
		if (m * m) < x {
			l = m + 1	
			res = m
		} else if (m * m) > x {
			r = m - 1
		} else {
			return m
		}
	}
	return res
}
