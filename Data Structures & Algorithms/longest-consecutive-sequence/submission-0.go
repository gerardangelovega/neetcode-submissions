func longestConsecutive(nums []int) int {
	m := make(map[int]bool) // map to store nums
	f := make(map[int]bool) // map to track consecutive
	max_consecutive := 0
	consecutive := 0

	for _, num := range nums {
		m[num] = true
	}

	for i := 0; i < len(nums); i++ {
		if f[i] {
			continue
		}

		if !m[nums[i] - 1] {
			num := nums[i]

			for m[num] {
				consecutive++
				num++
			}

			max_consecutive = max(max_consecutive, consecutive)
			consecutive = 0
		}
	}

	return max_consecutive
}
