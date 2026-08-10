func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }
	m := make(map[byte]int)
	for i := range len(s) {
		m[s[i]]++
		m[t[i]]--
	}
	for _, val := range m {
		if val != 0 { return false }
	}
	return true
}
