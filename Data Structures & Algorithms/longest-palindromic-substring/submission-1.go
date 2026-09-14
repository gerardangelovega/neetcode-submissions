func longestPalindrome(s string) string {
    res := ""
	n := len(s)
	for i := 0; i < n; i++ {
		l, r := i, i
		for {
			if len(s[l:r+1]) > len(res) {
				res = s[l:r+1]
			}
			l--
			r++
			if l < 0 || r > n-1 {
				break
			}
			if s[l] != s[r] {
				break
			}
		}
		if i - 1 > -1 {
			l, r = i - 1, i
			for {
				if s[l] != s[r] {
					break
				}
				if len(s[l:r+1]) > len(res) {
					res = s[l:r+1]
				}
				l--
				r++
				if l < 0 || r > n-1 {
					break
				}
			}
		}
		if i + 1 < n {
			l, r = i, i + 1
			for {
				if s[l] != s[r] {
					break
				}
				if len(s[l:r+1]) > len(res) {
					res = s[l:r+1]
				}
				l--
				r++
				if l < 0 || r > n-1 {
					break
				}
			}
		}
	}
	return res
}

func check(s string) bool {
	l, r := 0, len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}


// How do I handle even and odd length palindromes?
// I think for odd length:
// - I will start with a char at i and expand left and right until an invalid state
// - For every valid state, I will check if it is the longest palindrome
// I think for even length:
// - I will check if the char at i and the char before it are the same
// - Or, I will check the char at i and the char after it are the same
// - Then do the same as above