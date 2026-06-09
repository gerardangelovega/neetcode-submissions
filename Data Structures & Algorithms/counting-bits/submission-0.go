func countBits(n int) []int {
	arr := make([]int, n + 1)

	for i := range (n + 1) {
		arr[i] = count1Bits(i)
	}

	return arr
}

func count1Bits(n int) int {
	bits := 0

	for n != 0 {
		if n % 2 != 0 { bits++ }
		n = n >> 1
	}

	return bits
}