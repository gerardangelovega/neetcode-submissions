func anagramMappings(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, len(nums1))

	for _, num := range nums1 {
		m[num] = -1
	}

	for i, num := range nums2 {
		if _, exists := m[num]; exists {
			m[num] = i
		}
	}

	for i, num := range nums1 {
		res[i] = m[num]
	}

	return res
}
