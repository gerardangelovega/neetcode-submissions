func longestConsecutive(nums []int) int {
	m := make(map[int]bool) // map to store nums
	max_consecutive := 0

	for _, num := range nums {
		m[num] = true
	}

	for i := 0; i < len(nums); i++ {
		if !m[nums[i] - 1] {
			num := nums[i]
			consecutive := 0

			for m[num] {
				consecutive++
				num++
			}

			max_consecutive = max(max_consecutive, consecutive)
		}
	}

	return max_consecutive
}
