func uniquePathsWithObstacles(grid [][]int) int {
	return dfs(0, 0, len(grid), len(grid[0]), grid, make(map[[2]int]int))	
}

func dfs(r, c, m, n int, grid [][]int, cache map[[2]int]int) int {
	if r == m || c == n { return 0 }
	if grid[r][c] == 1 { return 0 }
	if r == (m - 1) && c == (n - 1) { return 1 }
	key := [2]int{r,c}
	if val, exists := cache[key]; exists {
		return val
	}
	cache[key] = dfs(r + 1, c, m, n, grid, cache) + dfs(r, c + 1, m, n, grid, cache)
	return cache[key]
}