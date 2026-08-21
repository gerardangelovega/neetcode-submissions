func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	for i, num := range nums {
		for j := range k {
			if i == min(n-1, i+j+1) { continue }
			if nums[min(n-1, i+j+1)] == num { return true }
		}
	}
	return false
}