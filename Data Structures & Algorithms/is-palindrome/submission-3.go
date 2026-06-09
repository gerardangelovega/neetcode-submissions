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

	reverse := ""

	for i := len(str) - 1; i >= 0; i = i - 1 {
		reverse = reverse + string(str[i])
	}

	return str == reverse
}
