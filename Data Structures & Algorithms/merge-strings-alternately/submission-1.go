func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	l, r, n1, n2 := 0, 0, len(word1), len(word2)
	for ; l < n1 && r < n2; l, r = l+1, r+1 {
		res.WriteByte(word1[l])
		res.WriteByte(word2[r])
	}
	for ; l < n1; l++ { res.WriteByte(word1[l]) }
	for ; r < n2; r++ { res.WriteByte(word2[r]) }
	return res.String()
}