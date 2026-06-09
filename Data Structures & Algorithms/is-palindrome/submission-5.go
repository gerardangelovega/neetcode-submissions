func isPalindrome(s string) bool {
	str := ""

	for i, r := range s {
		_ = i
		isalphanum := unicode.IsLetter(r) || unicode.IsNumber(r)
		if !isalphanum {
			continue
		}
		str = str + string(r)
	}

	str = strings.ToLower(str)

	for i, j := 0, len(str) - 1; i < len(str); i, j = i + 1, j - 1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}