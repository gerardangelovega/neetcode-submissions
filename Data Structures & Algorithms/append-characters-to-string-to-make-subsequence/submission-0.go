func appendCharacters(s string, t string) int {
	sn, tn := len(s), len(t)
	p1 := 0

	for i := 0; i < sn && p1 < tn; i++ {
		if s[i] == t[p1] {
			p1++
		}
	}

	return tn - p1
}