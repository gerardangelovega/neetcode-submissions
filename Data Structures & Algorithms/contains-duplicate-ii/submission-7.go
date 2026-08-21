func containsNearbyDuplicate(nums []int, k int) bool {
	dupes := make(map[int]int)
	for i, n := range nums {
		if j, e := dupes[n]; e && (i-j) <= k { return true }
		dupes[n] = i
	}
	return false
}
