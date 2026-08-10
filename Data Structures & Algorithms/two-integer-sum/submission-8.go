func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, exists := m[target-num]; exists {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}