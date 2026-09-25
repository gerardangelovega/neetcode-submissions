func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	l, r, n1, n2 := 0, 0, len(word1), len(word2)
	for l < n1 && r < n2 {
		res.WriteByte(word1[l])
		l++
		res.WriteByte(word2[r])
		r++
	}
	for l < n1 {
		res.WriteByte(word1[l])
		l++
	}
	for r < n2 {
		res.WriteByte(word2[r])
		r++
	}
	return res.String()
}
