func maxArea(heights []int) int {
	res := 0
	n := len(heights)

	for l, r := 0, n-1; l < r; {
		d := r - l
		lh, rh := heights[l], heights[r]
		h := min(lh, rh)
		res = max(res, h * d)

		if lh < rh {
			l++
		} else {
			r--
		}
	}

	return res
}