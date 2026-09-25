func mergeAlternately(word1 string, word2 string) string {
	l, r, n, m := 0, 0, len(word1), len(word2)
	var res strings.Builder
	res.Grow(n + m)
	for ; l < n && r < m; l, r = l+1, r+1 {
		res.WriteByte(word1[l])
		res.WriteByte(word2[r])
	}
	for ; l < n; l++ { res.WriteByte(word1[l]) }
	for ; r < m; r++ { res.WriteByte(word2[r]) }
	return res.String()
}