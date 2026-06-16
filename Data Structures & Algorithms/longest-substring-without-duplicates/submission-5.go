func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	l, r, n := 0, 0, len(s)
	res := 0
	
	for l < n && r < n {
		if _, exists := m[s[r]]; exists {
			delete(m, s[l])
			l++
		} else {
			m[s[r]] = struct{}{}
			r++
		}
		res = max(res, r - l)
	}

	return res
}
