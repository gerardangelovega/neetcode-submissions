func hammingWeight(n int) int {
	bits := 0
	num := n
	for i := range 32 {
		_ = i
		if num % 2 != 0 {
			bits++
		}
		num = num >> 1
	}

	return bits
}
