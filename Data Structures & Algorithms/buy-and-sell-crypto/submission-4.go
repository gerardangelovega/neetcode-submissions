func maxProfit(prices []int) int {
	max_profit := 0

	for i := 0; i < len(prices); i = i + 1 {
		for j := i + 1; j < len(prices); j = j + 1 {
			profit := prices[j] - prices[i]

			if profit > max_profit {
				max_profit = profit
			}
		}
	}

	return max(0, max_profit)
}