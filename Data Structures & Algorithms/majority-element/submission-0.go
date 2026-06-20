func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for key, value := range m {
		if value > (len(nums) / 2) {
			return key
		}
	}

	return -1
}
