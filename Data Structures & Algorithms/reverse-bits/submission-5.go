func reverseBits(n int) int {
	num := 0

	for i := 0; i < 32; i++ {
		if n & (1 << i) != 0 { num = num + 1 }
		num = num << 1
	}
	num = num >> 1

	return num
}
