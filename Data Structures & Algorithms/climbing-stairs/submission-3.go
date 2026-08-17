func climbStairs(n int) int {
	if n == 0 { return 0 }
    dp := [2]int{1, 2} 
    if n < 3 { return dp[n-1] }
	for i := 3; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0] + dp[1]
	}
	return dp[1]
}
