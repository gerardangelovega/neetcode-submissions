func reverseBits(n int) int {
	res := 0
	for range 31 {
		res = res + (n & 1)
		res = res << 1
		n = n >> 1
		fmt.Println(res)
	}
	res = res + (n & 1)
	return res
}
