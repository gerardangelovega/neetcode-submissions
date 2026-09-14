func longestPalindrome(s string) string {
    res := ""
	n := len(s)
	for i := 0; i < n; i++ {
		bfs(s, i, i, &res)	
		bfs(s, i - 1, i, &res)	
		bfs(s, i, i + 1, &res)	
	}
	return res
}

func bfs(s string, l, r int, res *string) {
	for l > -1 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		if len(s[l:r+1]) > len(*res) {
			*res = s[l:r+1]
		}
		l--
		r++	
	}
}