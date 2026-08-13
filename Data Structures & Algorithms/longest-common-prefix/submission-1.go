func longestCommonPrefix(strs []string) string {
	shortest := math.MaxInt
	for _, s := range strs {
		shortest = min(shortest, len(s))
	}

	res := ""
	for i := range shortest {
		for _, s := range strs[1:] {
			if s[i] != strs[0][i] {
				return string(res)
			}
		}
		res = res + string(strs[0][i])
	}
	return res
}
