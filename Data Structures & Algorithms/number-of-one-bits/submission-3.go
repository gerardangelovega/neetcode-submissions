func hammingWeight(n int) int {
	res := 0
	for b := n; b != 0; b = b >> 1 {
		if b & 1 == 1 { res++ }
	}
	return res
}
