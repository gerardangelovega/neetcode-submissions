func hasDuplicate(nums []int) bool {
	res := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := res[num]; !exists {
			res[num] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
