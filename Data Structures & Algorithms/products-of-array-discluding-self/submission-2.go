func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	pre := make([]int, n)
	suf := make([]int, n)

	pre[0], suf[n-1] = 1, 1
	for i, j := 1, n-2; i < n && j >= 0; i, j = i+1, j-1 {
		pre[i] = nums[i-1] * pre[i-1]
		suf[j] = nums[j+1] * suf[j+1]
	}

	for i := 0; i < n; i++ {
		res[i] = pre[i] * suf[i]
	}

	return res
}