func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pre := make([]int, len(nums))
	suf := make([]int, len(nums))

	for i, j := 0, len(nums) - 1; i < len(nums) && j >= 0; i, j = i + 1, j - 1 {
		if i == 0 && j == len(nums) - 1 {
			pre[i] = 1
			suf[j] = 1
			continue
		}

		pre[i] = nums[i - 1] * pre[i - 1]
		suf[j] = nums[j + 1] * suf[j + 1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = pre[i] * suf[i]
	}

	return res
}
