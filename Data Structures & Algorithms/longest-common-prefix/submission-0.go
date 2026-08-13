func longestCommonPrefix(strs []string) string {
	shortest := math.MaxInt
	for _, s := range strs {
		shortest = min(shortest, len(s))
	}

	res := ""
	for i := range shortest {
		curr := strs[0][i]
		fmt.Println(curr)
		same := true
		for _, s := range strs[1:] {
			if s[i] != curr {
				same = false
				break
			}
		}
		if same {
			res = res + string(curr)
		} else {
			break
		}
	}
	return res
}
