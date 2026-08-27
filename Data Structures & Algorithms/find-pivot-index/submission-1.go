func pivotIndex(nums []int) int {
	n := len(nums)
	pre, pos := make([]int, n), make([]int, n)
	copy(pre, nums)
	copy(pos, nums)
	for i := 1; i < n; i++ { pre[i] = pre[i-1] + pre[i] }
	for i := n-2; i >= 0; i-- { pos[i] = pos[i] + pos[i+1] }
	for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
		if 	pre[l] == pos[l] { return l }
		if 	pre[r] == pos[r] { return r }
	}
	return -1
}