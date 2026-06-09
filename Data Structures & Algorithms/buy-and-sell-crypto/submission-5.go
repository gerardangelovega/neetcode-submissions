func maxProfit(prices []int) int {
	max_profit := 0

	for i, j := 0, 1; j < len(prices); j = j + 1 {
		if prices[j] > prices[i] {
			max_profit = max(max_profit, prices[j] - prices[i])
		} else {
			i = j
		}
	}

	return max_profit
}
