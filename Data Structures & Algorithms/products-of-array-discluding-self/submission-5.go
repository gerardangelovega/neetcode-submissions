func productExceptSelf(nums []int) []int {
	n := len(nums)
	res, pre, pos := make([]int, n), make([]int, n), make([]int, n)
	pre[0], pos[n-1] = 1, 1
	for l, r := 1, n-2; l < n && r >= 0; l, r = l+1, r-1 {
		pre[l] = pre[l-1] * nums[l-1]
		pos[r] = pos[r+1] * nums[r+1]
	}
	for i, _ := range res { res[i] = pre[i] * pos[i] }
	return res
}
