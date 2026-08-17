func uniquePaths(m int, n int) int {
	return memoize(0, 0, m, n, make(map[[2]int]int))
}

func memoize(r, c, m, n int, cache map[[2]int]int) int {
	if r == m || c == n {
		return 0
	}
	if r == (m - 1) && c == (n - 1) {
		return 1
	}
	if val, exists := cache[[2]int{r,c}]; exists {
		return val
	}
	cache[[2]int{r,c}] = memoize(r + 1, c, m, n, cache) + memoize(r, c + 1, m, n, cache)
	return cache[[2]int{r,c}]
}