func reverseBits(n int) int {
	num := 0

	for i := 0; i < 31; i++ {
		if n & 1 != 0 {
			num = num + 1
			num = num << 1
		} else {
			num = num << 1
		}
		n = n >> 1
	}

	if n & 1 != 0 {
		num = num + 1
	} 

	return num
}
