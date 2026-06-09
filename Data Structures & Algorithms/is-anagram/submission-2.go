func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for i := 0; i < len(s); i = i + 1 {
		map1[rune(s[i])]++
		map2[rune(t[i])]++
	}

	for i := 0; i < len(s); i = i + 1 {
		if map1[rune(s[i])] != map2[rune(s[i])] {
			return false
		}
	}

	return true
}