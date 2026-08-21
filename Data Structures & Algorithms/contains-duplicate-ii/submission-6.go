func containsNearbyDuplicate(nums []int, k int) bool {
	dupes := make(map[int]struct{})
	l, r, n := 0, 0, len(nums)
	for r < n {
		if (r - l + 1) > k+1 { 
			delete(dupes, nums[l]) 
			l++
		}
		if _, e := dupes[nums[r]]; e { 
			return true 
		}
		dupes[nums[r]] = struct{}{}
		r++
	}
	return false
}
