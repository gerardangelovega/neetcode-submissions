func reverseString(s []byte) {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
