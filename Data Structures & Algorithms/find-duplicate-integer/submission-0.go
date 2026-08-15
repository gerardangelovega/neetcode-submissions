func findDuplicate(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, e := m[num]; e { return num }
		m[num] = struct{}{}
	}
	return -1
}