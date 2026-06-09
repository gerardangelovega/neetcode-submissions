func hammingWeight(n int) int {
	bits := 0
	num := n
	for num != 0 {
		if num % 2 != 0 {
			bits++
		}
		num = num >> 1
	}

	return bits
}
