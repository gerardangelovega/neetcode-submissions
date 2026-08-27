func pivotIndex(nums []int) int {
	n := len(nums)
	pre := make([]int, n)
	pos := make([]int, n)
	copy(pre, nums)
	copy(pos, nums)

	for i := 1; i < n; i++ { pre[i] = pre[i-1] + pre[i] }
	for i := n-2; i >= 0; i-- { pos[i] = pos[i] + pos[i+1] }

	l, r := 0, 0
	for l < n && r < n {
		if pre[l] == pos[r] { return l }
		l++
		r++
	}

	return -1
}
