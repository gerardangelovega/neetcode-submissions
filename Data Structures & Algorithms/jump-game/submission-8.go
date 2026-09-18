func canJump(nums []int) bool {
	n := len(nums)
	reach := 0

	for i := 0; i < n; i++ {
		if i > reach {
			return false
		}
		if reach >= n-1 {
			return true
		}
		reach = max(reach, i + nums[i])
	}
	return false
}
