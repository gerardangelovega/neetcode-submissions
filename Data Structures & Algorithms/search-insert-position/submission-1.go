func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r, m := 0, n-1, 0
	closest := 0

	if target < nums[l] { return 0 }
	if target > nums[r] { return n }

	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			if abs(nums[m]-target) < abs(nums[closest]-target) {
				closest = m
			}
			r = m-1
		} else {
			if abs(nums[m]-target) < abs(nums[closest]-target) {
				closest = m
			}
			l = m+1
		}
	}
	fmt.Println(closest, nums[closest])
	if nums[closest] > target {
		return closest
	}
	return closest+1
}

func abs(x int) int {
	if x < 0 { x = x * -1}
	return x
}