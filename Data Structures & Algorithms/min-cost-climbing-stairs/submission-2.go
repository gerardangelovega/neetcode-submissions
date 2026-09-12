func minCostClimbingStairs(cost []int) int {
	cache := make([]int, len(cost))
	for i, _ := range cache {
		cache[i] = -1
	}
	return min(climb(cost, 0, cache), climb(cost, 1, cache))
}

func climb(cost []int, i int, cache []int) int {
	if i >= len(cost) {
		return 0
	}
	if cache[i] != -1 {
		return cache[i]
	}
	cache[i] = cost[i] + min(climb(cost, i + 1, cache), climb(cost, i + 2, cache))

	return cache[i]
}

// 1, 2
// 2, 3, 1, 3+