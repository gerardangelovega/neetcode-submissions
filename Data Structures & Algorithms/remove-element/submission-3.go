func removeElement(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
			continue
		}
		i++
	}

	return n
}
