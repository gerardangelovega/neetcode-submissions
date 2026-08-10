func search(nums []int, target int) int {
	n := len(nums)
	l, r, m := 0, n - 1, (n - 1) / 2

	for l <= r {
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1	
		} else {
			l = m + 1
		}
		m = (l + r) /  2
	}

	return -1
}
