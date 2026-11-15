func removeVowels(s string) string {
	vowels := map[rune]struct{}{
		rune('a'): struct{}{},
		rune('e'): struct{}{},
		rune('i'): struct{}{},
		rune('o'): struct{}{},
		rune('u'): struct{}{},
	}
	res := ""
	for _, c := range s {
		if _, e := vowels[c]; e { continue }
		res = res + string(c)
	}
    return res
}
