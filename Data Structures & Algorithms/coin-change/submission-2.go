func coinChange(coins []int, amount int) int {
	res := dfs(coins, amount, make(map[int]int))
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func dfs(coins []int, amount int, cache map[int]int) int {
	if amount == 0 {
		return 0
	}
	if v, e := cache[amount]; e {
		return v
	}
	res := math.MaxInt32
	for _, c := range coins {
		if (amount - c) > -1 {
			res = min(res, 1 + dfs(coins, amount-c, cache))
		}
	}
	cache[amount] = res
	return res
}