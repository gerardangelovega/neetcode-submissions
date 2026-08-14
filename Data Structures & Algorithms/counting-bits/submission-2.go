func countBits(n int) []int {
	res := make([]int, n+1)
	for i := range n+1 {
		for b := i; b != 0; b = b >> 1 {
			if b & 1 == 1 { res[i]++ }
		}
	}
	return res
}