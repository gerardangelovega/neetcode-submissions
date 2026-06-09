func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	seq := make([]int, n + 1)
	seq[1] = 1
	seq[2] = 2

	for i := 3; i <= n; i++ {
		seq[i] = seq[i - 1] + seq[i - 2]
	}

	return seq[n]
}
