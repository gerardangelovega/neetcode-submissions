func mySqrt(x int) int {
	res, diff := 0, math.MaxInt

	l, r, m := 0, x, 0
	for l <= r {
		m = (l + r) / 2
		d := delta(x, m * m)

		if d > 0 {
			l = m + 1	
		} else if d < 0 {
			r = m - 1
		} else {
			return m
		}

		if abs(d) < abs(diff) && (m * m) <= x {
			diff = d
			res = m
		}
	}
	return res
}

func delta(a, b int) int {
	fmt.Println(a, b, a-b)
	return a - b
}

func abs(x int) int {
	if x < 0 { x = x * -1 }
	return x
}