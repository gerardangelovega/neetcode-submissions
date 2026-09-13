func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	c1, c2 := make(map[int]int), make(map[int]int)
    return max(dfs(nums[:n-1], 0, c1), dfs(nums[1:], 0, c2))
}

func dfs(nums []int, i int, cache map[int]int) int {
	if i >= len(nums) {
		return 0
	}
	if val, e := cache[i]; e {
		return val
	}
	cache[i] = max(nums[i] + dfs(nums, i + 2, cache), dfs(nums, i + 1, cache))
	return cache[i]
}