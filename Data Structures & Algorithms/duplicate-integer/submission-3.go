func hasDuplicate(nums []int) bool {
	m := make(map[int]bool, 100_000)

	for i := 0; i < len(nums); i = i + 1 {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
