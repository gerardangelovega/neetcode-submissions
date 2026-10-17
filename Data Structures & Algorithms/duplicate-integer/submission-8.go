func hasDuplicate(nums []int) bool {
	res := make(map[int]struct{}, len(nums) * 4)
	for _, n := range nums { if _, e := res[n]; e { return true } else { res[n] = struct{}{} } }
	return false
}
