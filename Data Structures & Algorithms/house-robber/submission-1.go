func rob(nums []int) int {
	return dfs(nums, make(map[int]int), 0)
}

func dfs(nums []int, cache map[int]int, n int) int {
	if n >= len(nums) {
		return 0
	}
	if val, e := cache[n]; e {
		return val
	}
	cache[n] = max(nums[n] + dfs(nums, cache, n + 2), dfs(nums, cache, n + 1))
	return cache[n]
}