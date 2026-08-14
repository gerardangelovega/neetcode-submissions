func hammingWeight(n int) int {
	res := 0
	for ; n != 0; n = n >> 1 {
		if n & 1 == 1 { res++ }
	}
	return res
}
