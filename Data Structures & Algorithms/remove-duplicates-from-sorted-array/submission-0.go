func removeDuplicates(nums []int) int {
	l, r, n := 0, 0, len(nums)
	for r < n {
		nums[l]	= nums[r]
		for r < n && nums[r] == nums[l] { r++ }
		l++
	}
	return l
}
