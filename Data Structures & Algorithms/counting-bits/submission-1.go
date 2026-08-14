func countBits(n int) []int {
	res := []int{}
	for i := range n+1 {
		bits := 0
		for ; i != 0; i = i >> 1 {
			if i & 1 == 1 { bits++ }
		}
		res = append(res, bits)
	}
	return res
}