func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, r := range s {
		map1[r] = map1[r] + 1
	}

	for _, r := range t {
		map2[r] = map2[r] + 1
	}

	for _, r := range s {
		if map1[r] != map2[r] {
			return false
		}
	}

	return true
}
