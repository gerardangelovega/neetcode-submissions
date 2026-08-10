func twoSum(nums []int, target int) []int {
    m := make(map[int32]int16)
	for i, num := range nums {
		if j, exists := m[int32(target-num)]; exists {
			return []int{int(j), i}
		}
		m[int32(num)] = int16(i)
	}
	return []int{}
}
