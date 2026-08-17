func climbStairs(n int) int {
	if n < 3 { return n }
    dp := [2]int{1, 2} 
	for i := 3; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0] + dp[1]
	}
	return dp[1]
}
