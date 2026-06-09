func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := make(map[rune]bool)
	unique := 0
	max_sub := 0

	for _, r := range s {
		m[r] = true
	}
	unique = len(m)
	clear(m)

	for i := unique; i > 0; i-- {
		for j := 0; j + i - 1 < n; j++ {
			for _, r := range s[j:(j+i)] {
				if m[r] {
					break
				}
				m[r] = true
			}
			max_sub = max(max_sub, len(m))
			clear(m)
		}
	}

	return max_sub
}
