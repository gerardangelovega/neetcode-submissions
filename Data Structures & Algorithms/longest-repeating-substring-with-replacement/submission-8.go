func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	l, r, n := 0, 1, len(s)

	m[s[0]]++
	maxfreq := m[s[0]]
	res := maxfreq

	for r < n {
		m[s[r]]++
		maxfreq = max(maxfreq, m[s[r]])
		for (r-l+1) - maxfreq > k {
			m[s[l]]--
			l++
			if m[s[l]] > maxfreq {
				maxfreq = m[s[l]]
			}
		}
		res = max(res, r-l+1)
		r++
	}

	return res
}
